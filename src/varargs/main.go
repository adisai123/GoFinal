package main

import (
	"fmt"
)

func main() {
	varar(10, false)
	varar(10)
}

func varar(a int, bo ...bool) {
	fmt.Println("ad", bo)
	fmt.Println("ad1", a)
}
