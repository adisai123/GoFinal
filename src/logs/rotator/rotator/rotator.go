package rotator

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// nl is a byte slice containing a newline byte.  It is used to avoid creating
// additional allocations when writing newlines to the log file.
var nl = []byte{'\n'}

// A Rotator writes input to a file, splitting it up into gzipped chunks once
// the filesize reaches a certain threshold.
type Rotator struct {
	size      int64
	threshold int64
	maxRolls  int
	filename  string
	out       *os.File
	wg        sync.WaitGroup
	mu        sync.RWMutex
}

// New returns a new Rotator.  The rotator can be used either by reading input
// from an io.Reader by calling Run, or writing directly to the Rotator with
// Write.
func New(filename string, thresholdKB int64, maxRolls int) (*Rotator, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("hd", filepath.Dir(filename))
		fmt.Println("hd1", err)
		os.Mkdir(filepath.Dir(filename), 777)
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	fmt.Println("hd2", err)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &Rotator{
		size:      stat.Size(),
		threshold: 1000 * thresholdKB,
		maxRolls:  maxRolls,
		filename:  filename,
		out:       f,
	}, nil
}

// Write implements the io.Writer interface for Rotator.  If p ends in a newline
// and the file has exceeded the threshold size, the file is rotated.
func (r *Rotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	n, _ = r.out.Write(p)
	r.size += int64(n)

	if r.size >= r.threshold && len(p) > 0 && p[len(p)-1] == '\n' {
		err := r.rotate()
		if err != nil {
			return 0, err
		}
		r.size = 0
	}
	r.mu.Unlock()
	return n, nil
}

// Close closes the output logfile.
func (r *Rotator) Close() error {
	err := r.out.Close()
	r.wg.Wait()
	return err
}

func (r *Rotator) rotate() error {
	dir := filepath.Dir(r.filename)
	glob := filepath.Join(dir, filepath.Base(r.filename)+".*")
	existing, err := filepath.Glob(glob)
	if err != nil {
		return err
	}

	maxNum := 0
	for _, name := range existing {
		parts := strings.Split(name, ".")
		if len(parts) < 2 {
			continue
		}
		numIdx := len(parts) - 1
		if parts[numIdx] == "gz" {
			numIdx--
		}
		num, err := strconv.Atoi(parts[numIdx])
		if err != nil {
			continue
		}
		if num > maxNum {
			maxNum = num
		}
	}

	err = r.out.Close()
	if err != nil {
		return err
	}
	str := time.Now().Format("20060102150405")
	rotname := fmt.Sprintf("%s.%s.%d", r.filename, str, maxNum+1)
	err = os.Rename(r.filename, rotname)
	if err != nil {
		return err
	}
	if r.maxRolls > 0 {
		for n := maxNum + 1 - r.maxRolls; n >= 0; n-- {
			files1, _ := filepath.Glob(r.filename + ".*." + strconv.Itoa(n))
			files2, _ := filepath.Glob(r.filename + ".*." + strconv.Itoa(n) + ".gz")
			file := append(files1, files2...)
			for _, f := range file {
				err := os.Remove(f)
				if err != nil {
					break
				}
			}
		}
	}
	r.out, err = os.OpenFile(r.filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	compress(rotname)
	os.Remove(rotname)
	return nil
}

func compress(name string) (err error) {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	arc, err := os.OpenFile(name+".gz", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	z := gzip.NewWriter(arc)
	if _, err = io.Copy(z, f); err != nil {
		return err
	}
	if err = z.Close(); err != nil {
		return err
	}
	return arc.Close()
}
