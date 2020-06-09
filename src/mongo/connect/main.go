package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	ss, err := mgo.Dial("mongodb://localhost")
	fmt.Println("errr", err)
	fmt.Println("sess", ss)
}
