package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("name:", p.first, "age:", p.age)
}
func main() {

	p1 := person{
		first: "Ma",
		last:  "Az",
		age:   15,
	}

	p1.speak()
}
