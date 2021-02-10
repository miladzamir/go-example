package main

import "fmt"

func main() {
	//ananymos struc
	h1 := struct {
		firstName string
		lastName  string
		code      int
	}{
		firstName: "Milad",
		lastName:  "Dehghan",
		code:      442777995,
	}
	fmt.Println(h1)

	fmt.Println(h1.firstName)
}
