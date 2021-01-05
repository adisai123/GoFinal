package osdiff

import (
	"fmt"
	"runtime"
)

func Hi() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("aditya.linx_Windows")
	case "darwin":
		fmt.Println("aditya.MAC operating system")
	case "linux":
		fmt.Println("aditya.Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
