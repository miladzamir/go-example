package main

import "fmt"

func main() {
	//using composite literal
	a := [5]int{10, 64, 58, 98, 12}
	fmt.Println(a)

	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}
	fmt.Printf("%T \n", a)

}
