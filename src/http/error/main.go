package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/err", func(rs http.ResponseWriter, r *http.Request) {
		io.WriteString(rs, "hissiis")
	})
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		http.Error(rs, "hiss", 210)
	})
	http.ListenAndServe(":8080", nil)
}
