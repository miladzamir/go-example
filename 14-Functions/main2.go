package main

import "fmt"

func main() {
	//variadic parameters
	students("ali", "mamad", "hossein", "reza")

	x := sum(10, 20, 30, 40)
	fmt.Println(x)
}

func students(s ...string) {
	fmt.Println(s)

	for _, v := range s {
		fmt.Println("name:", v)
	}
}

func sum(x ...int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
