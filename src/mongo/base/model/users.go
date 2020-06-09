package model

type User struct {
	Name   string `json:"name"`
	Gender string `JSON:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}
