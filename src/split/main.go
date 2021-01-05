package main

import (
	"fmt"
	"strings"
)

type x struct {
	name string
}

func main() {
	t := x{}
	my(t)
}

func my(xx x) {
	s := strings.Split(xx.name, "..")
	if xx.name == "" {
		fmt.Println("heyyy")
	}
	v := len(s)
	if v == 2 {
		fmt.Println(s[1], v)
	}
}
