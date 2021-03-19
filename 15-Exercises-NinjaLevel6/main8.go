package main

import "fmt"

func main() {
	x := foo()

	fmt.Println(x())
}

func foo() func() int {
	return func() int {
		return 15
	}
}
