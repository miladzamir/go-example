package main

import "fmt"

const (
	a = iota + 1
	b
	c = iota
	d
)

func main() {
	//iota
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
