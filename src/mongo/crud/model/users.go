package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"id"`
}

func (u User) String() string {
	return fmt.Sprintf("Name:%s", u.Name)
}
