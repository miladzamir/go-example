package main

import "fmt"

func main() {
	//slice
	x := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T \n", x)
}
