package main

import "fmt"

func main() {
	// fmt.Println(fac(4))

	fmt.Println(fac2(4))
}

func fac(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fac(n-1)
}

func fac2(n int) int {
	lastValue := n
	for i := 1; i < n; i++ {
		lastValue = lastValue * (n - i)
	}
	return lastValue
}
