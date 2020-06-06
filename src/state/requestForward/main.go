package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Println("methods", r.Method)
		rs.Header().Set("location", "/forward")
		rs.WriteHeader(http.StatusSeeOther)

	})
	http.HandleFunc("/forward", func(rs http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rs, "heyyyy rammm")
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
