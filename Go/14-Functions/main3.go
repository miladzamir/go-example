package main

import "fmt"

func main() {
	xi := []int{10, 20, 30, 40}

	x, n := sum("myCalc is:", xi...)
	fmt.Println(x, n)
}

func sum(name string, x ...int) (string, int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return name, sum
}
