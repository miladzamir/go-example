package main

import "fmt"

var x int = 42
var y string = "Jaems Bound"
var z bool = true

func main() {
	s := fmt.Sprintf("type x: %v y: %s z: %v \n", x, y, z)

	println(s)
}
