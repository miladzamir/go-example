package main

import "fmt"

func main() {
	x := []int{10, 11, 12, 13, 14, 59, 5, 1, 16, 185}

	//Easy Programming :)))
	for i, v := range x {
		fmt.Println("index:", i, "value:", v)
	}
}
