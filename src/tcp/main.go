package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"os"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go request(conn)

		go response(conn)
	}
}

func init() {
	tpl, err := template.ParseFiles("a.txt")
	if err != nil {
		fmt.Println("There was an error parsing file", err)
	}

	friends := []string{"Alex", "Conor", "Ken", "Ronnie", "Patick", "Nina", "Jeremy", "Gentry", "Christian"}

	err = tpl.Execute(os.Stdout, friends)
	if err != nil {
		fmt.Println("There was an error executing template", err)
	}
}

func request(n net.Conn) {

	scanner := bufio.NewScanner(n)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

}

func response(n net.Conn) {

	fmt.Fprint(n, "HTTP/1.1 200 ok")
}
