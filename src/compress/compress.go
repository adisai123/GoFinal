package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	ar := os.Args
	compress(ar[1])
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
