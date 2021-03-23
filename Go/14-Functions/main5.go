package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	code      int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak(text string) {
	fmt.Println(s.person.firstName, s.person.lastName, "says:", text)
}

func main() {

	p1 := secretAgent{
		person: person{
			firstName: "Mohammad",
			lastName:  "alirezaii",
			code:      4422,
		},
		ltk: true,
	}

	fmt.Println(p1)

	p1.speak("salam")
}
