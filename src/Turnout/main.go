package main

import "fmt"

func main() {
	in1, in2, out1, out2 := make(chan int), make(chan int), make(chan int), make(chan int)
	b := make(chan bool, 2)
	go func() {
		for i := 0; i < 100; i += 2 {
			in1 <- i
		}
		close(in1)
	}()
	go func() {
		for i := 1; i < 100; i += 2 {
			in2 <- i
		}
		close(in2)
	}()
	go func(read1 <-chan int, read2 <-chan int, write1 chan<- int, write2 chan<- int) {
		var data int
		var more bool
		for {
			select {
			case data, more = <-read1:
			case data, more = <-read2:
			}
			if !more {
				fmt.Println("closing")
				close(write1)
				close(write2)
				break
			}
			select {
			case write1 <- data:
			case write2 <- data:

			}

		}
	}(in1, in2, out1, out2)
	go func(out1, out2 chan int) {
		var more1, more2 bool
		var data int
		for {
			select {
			case data, more1 = <-out1:
				fmt.Println("first", data)
			case data, more2 = <-out2:
				fmt.Println("second", data)
			}
			if !more1 && !more2 {
				b <- true
				b <- true
				break
			}

		}
	}(out1, out2)

	<-b
	<-b
}
