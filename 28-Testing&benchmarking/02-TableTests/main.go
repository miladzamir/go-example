package main

import "fmt"

func main() {
	fmt.Println(mySum(5,4))
}
func mySum(xi ...int) int{
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}