package main

import "fmt"

func main() {
	foo()
	myName("zamir")
}

func foo() {
	fmt.Println("yo im from foo function!")
}
func myName(n string) {
	fmt.Println("my name is ", n)
}
