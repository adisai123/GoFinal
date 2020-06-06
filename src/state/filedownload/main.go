package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		rs.Header().Set("Content-Type", "text/html; charset=utf-8")
		if r.Method == http.MethodPost {
			rs.Header().Set("Content-Type", "multipart/form-data; charset=utf-8")
			f, _, e := r.FormFile("file")
			//fmt.Fprintln(rs, f, e)
			if e != nil {
				http.Error(rs, e.Error(), http.StatusInternalServerError)
			}
			bytes, _ := ioutil.ReadAll(f)
			string := string(bytes)
			fmt.Fprintln(rs, string)
			//fmt.Fprintf(rs, "")
			//fmt.Fprintln(rs, "file header:", h)
		} else {
			io.WriteString(rs, `
		<form method="post" enctype="multipart/form-data">
		upload	<input type="file" name="file"/>
		<input type="submit">
		<br>
		</form>
		`)
		}
	})

	http.ListenAndServe(":8080", nil)
}
