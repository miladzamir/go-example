package main

import "fmt"

func main() {
	m := map[string]int{
		"reza":    25,
		"hossien": 28,
		"hichkas": 30,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
