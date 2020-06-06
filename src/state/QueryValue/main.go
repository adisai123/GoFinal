package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rs, "hi...", r.FormValue("q"))
	})
	http.ListenAndServe(":8080", nil)
}
