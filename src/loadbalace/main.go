package main

import "fmt"

func main() {
	ch := make(chan int)
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	b := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	go makeout(ch, ch1, ch2, b)
	go func() {
		for dat1 := range ch2 {
			fmt.Println("first", dat1)
		}
		b <- true
	}()
	go func() {
		for dat1 := range ch1 {
			fmt.Println("second", dat1)
		}
		b <- true
	}()
	<-b
	<-b
	<-b
	fmt.Println("end main")
}

func makeout(ch <-chan int, ch1, ch2 chan<- int, b chan<- bool) {
	for data := range ch {
		select {
		case ch1 <- data:
		case ch2 <- data:
		}

	}
	close(ch1)
	close(ch2)
	b <- true
	fmt.Println("end makeout")
}
