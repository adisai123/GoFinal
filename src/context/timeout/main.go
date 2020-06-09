package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctout, cancel := context.WithTimeout(ctx, 1*time.Second)
		ch := make(chan int, 1)
		defer cancel()
		go func() {
			time.Sleep(3 * time.Second)
			if ctx.Err() != nil {
				return
			}
			ch <- 10
		}()
		select {
		case <-ctout.Done():
			//log.Fatalln(ctout.Err())
			log.Println(ctout.Err())
		case i := <-ch:
			log.Println(i)
		}
	})
	http.ListenAndServe(":8080", nil)
}
