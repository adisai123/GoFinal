package main

import (
	"gofinal/src/src/mongo/base/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main()  {
	uc := controller.GetController()
	r := httprouter.New()
	r.GET("/",uc.Index)
	r.POST("/user/",uc.Create)
	r.GET("/users/:id",uc.GetUser)
	http.ListenAndServe(":8080",r)
}
