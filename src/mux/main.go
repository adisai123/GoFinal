package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func allusers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []user
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newuser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var users []user
	db.CreateTable(&users)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete user Endpoint Hit")
}

func updateuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allusers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteuser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateuser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newuser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// our initial migration function
func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&user{})
}

type user struct {
	firstName string
	email     string
}

func main() {
	fmt.Println("Go ORM Tutorial")

	// Handle Subsequent requests
	handleRequests()

}
