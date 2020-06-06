package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(rs http.ResponseWriter, r *http.Request) {
		rs.Header().Set("Content-Type", "text/html; charset=utf-8")
		q := r.FormValue("ison") == "on"
		io.WriteString(rs, `
		<form method="post">
		password: <input type="password" name="password"> 
		</br>
		check <input type="checkbox" name="ison">
		</br>
		<input type="submit" >
		</br>
		</form>
		`+strconv.FormatBool(q))
	})

	http.ListenAndServe(":8080", nil)
}
