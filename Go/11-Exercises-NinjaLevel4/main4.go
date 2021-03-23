package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	y := append(x, 52, 53, 54, 55)
	fmt.Println(y)

	n := []int{56, 57, 58, 59, 60}

	n = append(n, y...)

	fmt.Println(n)
}
