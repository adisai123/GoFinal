package main

import (
	"context"
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/",func(rs http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		str := r.FormValue("q")
		fmt.Println("",str)
		c := context.WithValue(ctx,"userID",str)
		var s string = c.Value("userID").(string)
		fmt.Fprintln(rs,s)
	})
	http.HandleFunc("/bar",func(rs http.ResponseWriter, r *http.Request){
		ctx := r.Context()
	
		fmt.Fprintln(rs,ctx.Value("userID"))
	})
	http.ListenAndServe(":8080",nil)
}