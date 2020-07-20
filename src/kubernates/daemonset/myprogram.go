package main

import (
	"fmt"
	"time"
)

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl

func doEvery(d time.Duration) {
	for x := range time.Tick(d) {
		fmt.Printf("%v: Hello, World!\n", x)

	}
}

func main() {
	doEvery(1 * time.Millisecond)
}
