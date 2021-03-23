package main

import "fmt"

const code int16 = 15161

const (
	a int = iota
	b
	c
)

func main() {
	fmt.Println(code)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
