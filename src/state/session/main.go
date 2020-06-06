package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", nil)
}
