package model

import (
	"log"

	"google.golang.org/genproto/googleapis/iam/v1/logging"
)

type Student struct{
	id string
	name string
	age int
}

func (s Student) String(){
	log.Printf("[id: %s, name: %s, age: %d ]\n",s.id,s.name,s.age)
}