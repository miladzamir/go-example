package main

import "fmt"

type human struct {
	firstName string
	lastName  string
	code      int
}
type secretAgent struct {
	human
	firstName string
	ltk       bool
}

func main() {

	h1 := human{
		firstName: "Milad",
		lastName:  "Dehghan",
		code:      442777995,
	}
	fmt.Println(h1)

	fmt.Println(h1.firstName)

	sa := secretAgent{
		human: human{
			firstName: "James",
			lastName:  "bound",
			code:      2400,
		},
		firstName: "SOMETIHNG",
		ltk:       true,
	}
	fmt.Println(sa)

	fmt.Println(sa.firstName, sa.lastName, sa.code, sa.ltk)

	//sa.human.firstName
	fmt.Println(sa.human.firstName, sa.lastName, sa.code, sa.ltk)
}
