package main

import "fmt"

func main() {
	x := make([]string, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = []string{"A", "new", "initialized", "slice", "value", "for", "a", "given", "element", "type", "T", "is", "made", "using", "the", "built-in", "function", "make"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println("index", i, "value", x[i])
	}
}
