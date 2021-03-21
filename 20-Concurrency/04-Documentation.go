package main

import "fmt"

func doSomething(x int) int {
	return 5 * x
}

func main() {
	ch := make(chan int)
	fmt.Println("Hello")
	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println("Hello2")
	fmt.Println(<-ch)

	fmt.Println("Hello3")
}
