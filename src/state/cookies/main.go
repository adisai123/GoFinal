package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/",func(rs http.ResponseWriter, r *http.Request){
		
		co,err:= r.Cookie("counter")
		if err == http.ErrNoCookie{
			co =&http.Cookie{
				Name: "counter",
				Value: "0",
			}
		}
		count, _ := strconv.Atoi(co.Value) 
		co.Value = strconv.Itoa(count + 1 ) 
		http.SetCookie(rs,co)
		fmt.Println("counter",co.Value)
		fmt.Fprintln(rs,"Written to clients browser....")
	})
	http.HandleFunc("/read",func(rs http.ResponseWriter, r *http.Request){
		co,err:= r.Cookie("counter")
		if err != nil {
			fmt.Fprintln(rs,err)
		}
		if co != nil {
			fmt.Fprintln(rs,co.MaxAge,co.Domain,co.Expires,co.Path)
			
		}
		fmt.Fprintln(rs,co)
			
	})
	http.HandleFunc("/exp",func(rs http.ResponseWriter, r *http.Request){
		co,err:= r.Cookie("counter")
		if err != nil {
			log.Fatalln(err)
		}
		co.MaxAge = -1   //delete cookies
		http.SetCookie(rs,co)
		fmt.Fprintln(rs,co)
			
	})

	http.ListenAndServe(":8080",nil)
}