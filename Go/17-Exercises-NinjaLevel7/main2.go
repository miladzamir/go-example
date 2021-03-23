package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Mohammad"
	// (*p).name = "Mohammad"
}

func main() {
	p1 := person{
		name: "Ali",
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
