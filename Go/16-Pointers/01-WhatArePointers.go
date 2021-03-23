package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Println(&a)
	fmt.Printf("%T \n", &a)

	var b *int = &a
	// b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 66

	fmt.Println(a)
}
