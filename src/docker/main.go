package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		rs.Write([]byte("this is container based go"))
	})
	http.ListenAndServe(":8080", nil)
}
