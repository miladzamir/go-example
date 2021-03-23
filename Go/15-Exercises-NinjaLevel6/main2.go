package main

import "fmt"

func main() {
	x := foo([]int{100, 500, 400, 500, 500}...)
	fmt.Println(x)

	y := bar([]int{100, 500, 400, 1000, 2000})
	fmt.Println(y)
}

func foo(n ...int) int {
	value := 0
	for _, v := range n {
		value += v
	}
	return value
}

func bar(n []int) int {
	value := 0
	for _, v := range n {
		value += v
	}
	return value
}
