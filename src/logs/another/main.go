package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/antigloss/go/logger"
)

var wg sync.WaitGroup

func main() {
	logger.Init("./log", 10, 2, 2, false)

	fmt.Print("Single goroutine (200000 writes), GOMAXPROCS(1): ")
	tSaved := time.Now()
	for i := 0; i != 200000; i++ {
		logger.Info("Failed to find player! uid=%d plid=%d cmd=%s xxx=%d", 1234, 678942, "getplayer", 102020101)
	}
	fmt.Println(time.Now().Sub(tSaved))

	fmt.Print("200000 goroutines (each makes 1 write), GOMAXPROCS(1): ")
	test()

	fmt.Print("200000 goroutines (each makes 1 write), GOMAXPROCS(2): ")
	runtime.GOMAXPROCS(2)
	test()

	fmt.Print("200000 goroutines (each makes 1 write), GOMAXPROCS(4): ")
	runtime.GOMAXPROCS(4)
	test()

	fmt.Print("200000 goroutines (each makes 1 write), GOMAXPROCS(8): ")
	runtime.GOMAXPROCS(8)
	test()
}

func test() {
	tSaved := time.Now()
	for i := 0; i != 200000; i++ {
		wg.Add(1)
		go func() {
			logger.Info("Failed to find player! uid=%d plid=%d cmd=%s xxx=%d", 1234, 678942, "getplayer", 102020101)
			wg.Add(-1)
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(tSaved))
}
