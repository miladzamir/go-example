package main

import "fmt"

func main() {
	//function expression

	f := func() {
		fmt.Println("yo!")
	}
	f()

	g := func(x, y int) (int, int) {
		return y, x
	}
	fmt.Println(g(5, 10))
}
