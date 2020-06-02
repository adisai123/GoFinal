package main

import (
	"io"
	"net/http"
)

func s(re http.ResponseWriter, r *http.Request) {
	io.WriteString(re, "about")
}

func c(re http.ResponseWriter, r *http.Request) {
	io.WriteString(re, "home")
}
func main() {
	var t T
	// HandlerFunc is a type an adaptor to allow the use of ordinary functions as HTTP handler
	http.Handle("/about", http.HandlerFunc(s))
	http.Handle("/my", t)
	http.HandleFunc("/home", c)
	http.HandleFunc("/", func(re http.ResponseWriter, r *http.Request) {
		io.WriteString(re, "inner")
	})
	http.ListenAndServe(":8080", nil) //if handler is null then add default serve mux , handlefunc
}

type T int

func (t T) ServeHTTP(re http.ResponseWriter, r *http.Request) {
	io.WriteString(re, "default")
}
