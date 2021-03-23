package main

import "fmt"

func main() {
	//delete value by key

	m := map[string]int{
		"reza":    25,
		"hossien": 28,
		"hichkas": 30,
	}
	fmt.Println(m)

	if v, ok := m["reza"]; ok {
		fmt.Println(v)
		delete(m, "reza")
	} else {
		fmt.Println("key not exsist")
	}

	fmt.Println(m)
}
