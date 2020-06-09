package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"log"
)

func main() {
	h := hmac.New(sha256.New, []byte("myjey"))
	io.WriteString(h, "aditya")
	log.Printf("%x", h.Sum(nil))

}