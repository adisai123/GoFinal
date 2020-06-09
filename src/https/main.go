package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		rs.Header().Set("Content-Type", "text/plain")
		rs.Write([]byte("hey ram"))
	})
	err := http.ListenAndServeTLS(":10344", "etcd.crt", "etcd.key", nil)
	fmt.Println(err)

}
