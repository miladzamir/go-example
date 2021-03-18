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

func (p person) speak(text string) {
	fmt.Println(p.firstName, p.lastName, "(person) says:", text)
}

//value can be of morethen one type
type human interface {
	speak(text string)
}

func bar(h human, text string) {
	switch h.(type) {
	case person:
		fmt.Println("person!", h.(person).firstName)
	case secretAgent:
		fmt.Println("secretAgent!", h.(secretAgent).ltk)
	}

	// fmt.Println("i called human \n", h)
	// h.speak(text)
}

func main() {

	sa := secretAgent{
		person: person{
			firstName: "Mohammad",
			lastName:  "alirezaii",
			code:      4422,
		},
		ltk: true,
	}

	p1 := person{
		firstName: "zam1r",
		lastName:  "alfred",
		code:      2545,
	}

	// fmt.Println(sa)

	// // sa.speak("salam")

	// fmt.Println(p1)

	bar(sa, "salam")
	bar(p1, "khodafez")
}
