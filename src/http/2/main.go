package main

import (
	"fmt"
	"net/http"
)

func main() {
	var m m
	http.ListenAndServe(":8080", m)
}

type m int

func (m m) ServeHTTP(rs http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/about":
		fmt.Fprintln(rs, "aboutus")
	default:
		fmt.Fprintln(rs, r.Method)
		rs.Header().Add("NUPUR", " UPUR")
		fmt.Fprintln(rs, rs.Header())

	}

}
