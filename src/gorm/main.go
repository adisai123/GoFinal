package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)
var myrouts *mux.Router 
var db *gorm.DB
func main()  {
	myrouts = mux.NewRouter().StrictSlash(true)
	handleRouts()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8080",myrouts))
	log.Println("ending")
}
type User struct {
	gorm.Model
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}
func (u User) String()string{
	return fmt.Sprintf("[name :%s , Gender: %s , Age: %d]",u.Name,u.Gender,u.Age)
}
func init()  {
	db, _ = gorm.Open("sqlite3","./foo1.db")

	db.AutoMigrate(&User{})
}
func handleRouts(){
	myrouts.HandleFunc("/",func(rs http.ResponseWriter,r *http.Request){
		rs.Write([]byte("starting point"))
	}).Methods("GET")
	myrouts.HandleFunc("/user",getAllUser).Methods("GET")
	myrouts.HandleFunc("/user/{Name}/{Gender}/{Age}",create).Methods("GET")
	myrouts.HandleFunc("/user/delete/{Name}",delete)
	myrouts.HandleFunc("/user/update/{Name}/{Gender}/{Age}",update)
}

func update(rs http.ResponseWriter, r *http.Request)  {
	vars := returnArgs(r)
	var use User
	fmt.Fprintln(rs,"Name",vars["Name"])
	fmt.Fprintln(rs,"Gender",vars["Gender"])
	fmt.Fprintln(rs,"Age",vars["Age"])
	db.Where("name = ?" , vars["Name"]).Find(&use)
	err := userFoundCheck(rs,use)
	if err != nil{
		use.Name = vars["Name"]
		rs.Write([]byte("\nNew record\n"))
	}
		use.Gender = vars["Gender"]
		use.Age,_ = strconv.Atoi(vars["Age"])
		db.Save(&use)
	fmt.Fprintln(rs,"record to update",use)
	rs.Write([]byte("Record updated"))
}
func userFoundCheck(rs http.ResponseWriter, use User) error{
	if use.Name == "" {
		rs.Write([]byte("User not found"))
		return errors.New("User Not found")
	}
	return nil
}
func delete(rs http.ResponseWriter, r *http.Request)  {
	args :=returnArgs(r)
	var usr User
	db.Where("name = ?" , args["Name"]).Find(&usr)
	err := userFoundCheck(rs,usr)
	if err == nil{
		db.Delete(&usr)
	}
	rs.Write([]byte("Element: "+args["Name"]+" New user created"))
}

func getAllUser(rs http.ResponseWriter, r *http.Request)  {
	var users []User
	db.Find(&users)
	log.Println(users)
	json.NewEncoder(rs).Encode(users)
}
func create(rs http.ResponseWriter, r *http.Request)  {
	vars := returnArgs(r)
	age,_ := strconv.Atoi(vars["Age"])
	db.Create(&User{Name:vars["Name"],Age:age,Gender:vars["Gender"]})
	rs.Write([]byte("New user created"))
}

func returnArgs(r *http.Request) map[string]string {
	return mux.Vars(r)
}