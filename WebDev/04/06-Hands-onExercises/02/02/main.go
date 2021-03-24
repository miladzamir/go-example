package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

		bufio.NewScanner(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn.RemoteAddr())
		io.WriteString(conn, "I see you connected \n")
		conn.Close()
	}
}
