package main

import "fmt"

var x int

//Own Type
type num float32

var b num

func main() {
	fmt.Printf("%T \n", x)
	fmt.Println(x)

	b = 500.600
	fmt.Printf("%T \n", b)
	fmt.Println(b)

	//Conversion NOT casting
	x = int(b)
	fmt.Printf("%T \n", x)
	fmt.Println(x)

}
