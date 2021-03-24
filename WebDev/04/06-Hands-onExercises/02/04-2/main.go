package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)
//Now I want you to READ from the connection.
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

			go handle(conn)
	}
}
func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Fprintf(conn, "Computer Says: %v \n",  ln)
		if ln == "" {
			break
			fmt.Println("END")
		}
	}
	defer conn.Close()

	io.WriteString(conn,"Connected")
}
