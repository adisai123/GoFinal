package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var logc = make(chan string, 1)

type str struct {
	s *sync.Mutex
}

//real    0m4.579s
var s *str = &str{}

func main() {
	for i := 0; i < 10000; i++ {
		func() {
			s.LogFileWrite(strconv.Itoa(i) + ": hi qsdjdqwkjd")
		}()
	}

}

func (s *str) LogFileWrite(tlog string) {
	f, _ := os.OpenFile("override.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	for {
		select {
		case logc <- tlog:
		}
		select {
		case log := <-logc:
			f.Truncate(0)
			f.Seek(0, 0)
			f.WriteString(log)
			fmt.Println(log)
		}
		break
	}
}
