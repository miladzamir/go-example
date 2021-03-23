package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Ali",
		Last:  "Mohammadi",
		Age:   90,
	}
	p2 := person{
		First: "Reaza",
		Last:  "Pishro",
		Age:   15,
	}
	p3 := person{
		First: "sos",
		Last:  "Mast",
		Age:   85,
	}

	person := []person{p1, p2, p3}

	fmt.Println(person)

	bs, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
