package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	f := x[:3]
	fmt.Println(f)

	s := x[6:]
	fmt.Println(s)

	l := append(f, s...)
	fmt.Println(l)
}
