package main

import "fmt"

func main()  {
	title := "Hello world!"

	tpl := `
			<!DOCTYPE html>
			<html>
			<head>
			<title>Page Title</title>
			</head>
			<body>
			<h1>`+ title +`</h1>
			</body>
			</html> 
			`
	fmt.Println(tpl)
}
/*

	go run main.go > index.html

*/