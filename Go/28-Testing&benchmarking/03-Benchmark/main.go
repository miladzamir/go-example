package main

import "fmt"

func main() {
	fmt.Println(calc(1,5))
}
func calc(n ...int) int{
	s := 0
	for _, n := range n {
		s += n
	}
	return s
}