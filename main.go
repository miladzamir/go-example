package main

import "fmt"

var x string = "b"
var y int

func main() {
	z := 110555

	fmt.Printf("%T \t %v %v \t \n", x, x, x)

	y = 1545516546

	fmt.Printf("%b \n", y)

	fmt.Println("my number is : ", z, "ok!")

	i := fmt.Sprintf("%s", x)

	fmt.Println(i, "\n")
}
