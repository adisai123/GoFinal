package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/all", func(rs http.ResponseWriter, r *http.Request) {
		rs.Header().Set("Content-Type", "text; charset=utf-8")
		f, err := os.Open("a.txt")
		if err != nil {
			panic(err)
		}
		b := make([]byte, 10000)
		f.Read(b)
		//rs.Write(b) different ways to add contents
		fmt.Fprint(rs, "<a href='./a.txt'>abcx<a>")
		//	i := io.MultiWriter(rs)
		http.ServeFile(rs, r, ".")

	})
	//if resources folder has index.html then it will load that only
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./resource"))))
	http.ListenAndServe(":8080", nil)
}
