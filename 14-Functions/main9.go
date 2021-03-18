package main

import "fmt"

func main() {

	fmt.Println(foo())

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

	//cleanup
	fmt.Println(bar()())
}
func foo() string {
	return "hello world!"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
