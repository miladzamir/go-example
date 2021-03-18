package main

import "fmt"

func main() {
	//map
	m := map[string]int{
		"reza": 25,
		// "hossien": 28,
		"hichkas": 30,
	}

	fmt.Println(m)

	fmt.Println("Reza age is :", m["reza"])

	if v, ok := m["hossien"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
	} else {
		fmt.Println(ok)
	}
}
