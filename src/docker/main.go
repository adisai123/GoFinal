package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var i uint32

func main() {
	rand.Seed(time.Now().UnixNano())
	i = rand.Uint32()
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Println("methods", r.Method)
		fmt.Println("request coming from :== ", i)
		rs.Header().Set("location", "/forward")
		rs.WriteHeader(http.StatusSeeOther)

	})
	http.HandleFunc("/forward", func(rs http.ResponseWriter, r *http.Request) {
		for i := 0; i < 100; i++ {
			fmt.Fprintln(rs, " !!!!!!🥰")
		}
	})
	http.HandleFunc("/red", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Println("red", r.Method)
		http.Redirect(rs, r, "/redd", http.StatusPermanentRedirect)
	})
	http.HandleFunc("/redd", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rs, "redd")

	})
	http.ListenAndServe(":8080", nil)
}
