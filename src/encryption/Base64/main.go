package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// encString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	s := base64.StdEncoding.EncodeToString([]byte(`HI I AM AD !@#$%^&*(()_{}|:"<>?~'`))
	fmt.Println(s)
	d, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(d))
}
