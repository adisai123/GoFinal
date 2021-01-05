package main

import (
	"fmt"

	"github.com/osdiff/osdiff"
)

func main() {
	var i uint8
	i = 0
	h := fmt.Sprintf("chanfge %x", int(i)<<2)
	fmt.Println(h)
	osdiff.Hi()
}
