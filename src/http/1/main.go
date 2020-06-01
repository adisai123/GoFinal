package main

import (
	"fmt"
	"net/http"
)

func main() {
	var i i = 10
	http.ListenAndServe(":8080", i)
}

type i int

//handler interface for http
func (i i) ServeHTTP(rs http.ResponseWriter, r *http.Request) {
	//rs.Write([]byte("hi"))
	fmt.Fprintln(rs, "another way to responde")
}
