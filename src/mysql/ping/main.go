package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	s, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/aditya")
	defer s.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Ping())
	row, err := s.Query("select id from abc")
	i := 0
	for row.Next() {
		row.Scan(&i)
		fmt.Println(i)
	}
}
