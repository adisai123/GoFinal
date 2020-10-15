package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		rs.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(rs, `
		<form method="post" enctype="multipart/form-data">
		upload	<input type="file" name="file"/>
		<input type="submit">
		<br>
		</form>
		`)
		if r.Method == http.MethodPost {
			f, h, e := r.FormFile("file")
			defer f.Close()
			//fmt.Fprintln(rs, f, e)
			if e != nil {
				http.Error(rs, e.Error(), http.StatusInternalServerError)
			}
			bytes, _ := ioutil.ReadAll(f)
			//string := string(bytes)
			//fmt.Fprintln(rs, string)
			file, _ := os.Create(filepath.Join("./", h.Filename))
			file.Write(bytes)
			//fmt.Fprintf(rs, "")
			//fmt.Fprintln(rs, "file header:", h)
		}

	})

	http.ListenAndServe(":8080", nil)
}
