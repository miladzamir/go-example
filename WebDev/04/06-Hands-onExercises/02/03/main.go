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
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		io.WriteString(conn, "Connected")
		conn.Close()
	}
}
