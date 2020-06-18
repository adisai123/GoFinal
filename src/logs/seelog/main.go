package main

import "gofinal/src/src/logs/seelog/config"

func main() {
	for i := 0; i < 500; i++ {
		go func(i int) {
			config.Logger.Critical("V2 . This is me ", i)
		}(i)

	}
}
