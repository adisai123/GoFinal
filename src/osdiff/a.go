package osdiff

import (
	"fmt"
	"runtime"
)

func Hi() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("nupurs.mac_Windows")
	case "darwin":
		fmt.Println("nupur.MAC operating system")
	case "linux":
		fmt.Println("nupur.Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
