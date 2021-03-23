package main

import "fmt"

type zamir int

var x zamir

func main() {

	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)

	x = 42

	fmt.Println(x)

}
