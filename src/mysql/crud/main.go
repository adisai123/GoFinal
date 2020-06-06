package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func init() {
	c, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/aditya")
	if err != nil {
		panic(err)
	}
	conn = c

}

func main() {
	createTable(conn)
	insert("aditya", conn)
	display(conn)
	defer conn.Close()
}

func display(conn *sql.DB) {
	rows, _ := conn.Query("select * from student")
	var name string
	var id int
	log.Println("-------------------select * from student----------------result:")
	log.Println("id\t\tname")
	log.Println("-----------------------------")
	for rows.Next() {
		rows.Scan(&name, &id)
		log.Println(id, "\t \t", name)
	}
}

func createTable(conn *sql.DB) {
	stmt, err := conn.Prepare("Create table student(name varchar(100), id int auto_increment, primary key(id)) ")
	if err != nil {
		panic(err)
	}
	rs, _ := stmt.Exec()
	if rs != nil {
		n, err := rs.RowsAffected()
		if err != nil {
			panic(err)
		}
		log.Println("number of tables get created", n)
		return
	}
	log.Println("Table already exist")

}

func insert(name string, conn *sql.DB) {
	query := "insert into student (name) values('" + name + "')"
	log.Println("query", query)
	stmt, err := conn.Prepare(query)
	if err != nil {
		panic(err)
	}
	rs, _ := stmt.Exec()
	if rs != nil {
		id, _ := rs.LastInsertId()
		log.Println("last inserted id", id)
	}
}
