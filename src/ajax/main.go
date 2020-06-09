package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil)
}

func foo(rs http.ResponseWriter, r *http.Request) {
	rs.Write([]byte("hiiii people"))
}
func index(rs http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(rs, "index.gohtml", nil)
}
