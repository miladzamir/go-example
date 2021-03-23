package main

import "fmt"

type human struct {
	firstName string
	lastName  string
	code      int
}

func main() {

	h1 := human{
		firstName: "Milad",
		lastName:  "Dehghan",
		code:      442777995,
	}
	fmt.Println(h1)

	fmt.Println(h1.firstName)

}
