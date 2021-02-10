package main

import "fmt"

type hotdog int

var x hotdog
var a int = 5
var y int

func main() {
	y = 200
	x = 100

	y = int(x)

	fmt.Println(y)

}
