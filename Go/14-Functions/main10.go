package main

import "fmt"

func main() {
	// fmt.Println(sum([]int{10, 50, 60}...))

	xi := evenNumbers(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	fmt.Println(xi)

	xii := oddNumbers(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(xii)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func evenNumbers(f func(x ...int) int, j ...int) int {
	var y []int
	for _, v := range j {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	total := f(y...)
	return total
}

func oddNumbers(f func(x ...int) int, j ...int) int {
	var y []int
	for _, v := range j {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
