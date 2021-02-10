package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main()  {
	name := "Sara"
	tpl := fmt.Sprint(`
			<!DOCTYPE html>
			<html>
			<head>
			<title>Page Title</title>
			</head>
			<body>
			<h6>my name is:`+ name +`</h6>
			</body>
			</html> 
			`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
/*

	go run main.go

*/