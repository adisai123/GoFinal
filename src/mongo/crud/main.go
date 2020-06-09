package main

import (
	"gofinal/src/src/mongo/crud/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	c := controller.GetMongoController(getSession())
	r.POST("/create/", c.Create)
	r.GET("/user/:id", c.Get)
	r.PUT("/user/update", c.Update)
	r.DELETE("/user/delete/:id", c.Delete)
	r.GET("/", c.Index)
	http.ListenAndServe(":8080", r)

}

func getSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}
