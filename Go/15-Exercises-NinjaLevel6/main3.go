package main

import "fmt"

func main() {

	defer foo()
	bar()

}

func foo() {
	fmt.Println("i'm first")
}

func bar() {
	fmt.Println("i'm second")
}
