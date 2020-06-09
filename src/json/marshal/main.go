package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type abc struct {
		Name string
		Age  int
	}
	aa := abc{Name: "aditya", Age: 100}
	fmt.Println(aa)
	bt, err := json.Marshal(aa)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bt))
	var a abc
	json.Unmarshal(bt, &a)
	fmt.Println(a)
}
