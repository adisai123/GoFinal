package main

import (
	"fmt"
	"gofinal/src/src/logs/rotator/rotator"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

//var err = os.MkdirAll("./Transaction", 0777)
var ro, _ = rotator.New("./Transaction/Transactions.log", 100, 10)

func main() {
	//defer ro.Close()
	var v sync.WaitGroup
	v.Add(5000)
	for i := 0; i < 5000; i++ {
		go func(i int) {

			ro.Write([]byte("V4 . This is me " + strconv.Itoa(i) + "\n"))
			v.Done()
		}(i)

	}
	v.Wait()
}

func done() {
	str := time.Now().Format("20060102150405")
	fmt.Println(time.Now())
	fmt.Println(str)
	filename := "./Transaction/" + str + "Transaction.log12"
	os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0777)
	myfile, _ := filepath.Glob("./Transaction/*Transaction.log12")
	for _, f := range myfile {
		fmt.Println(f)
		os.Remove(f)
	}
}
