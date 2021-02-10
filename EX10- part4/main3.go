package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 99}
	fmt.Println(x[:5])
	fmt.Println(x[4:9])
	fmt.Println(x[5:])
	fmt.Println(x[3:8])
}
