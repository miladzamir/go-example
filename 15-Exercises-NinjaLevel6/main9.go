package main

import "fmt"

func main() {
	x := foo(add, []int{15, 16, 185}...)
	fmt.Println(x)
}

func foo(f func(sum int) int, n ...int) int {

	sum := 0
	for _, v := range n {
		sum += v
	}
	return f(sum)
}

func add(sum int) int {
	return sum * 2
}
