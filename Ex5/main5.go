package main

import "fmt"

type zam1r int

var x zam1r
var y int

func main() {
	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)

	x = 42

	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
}
