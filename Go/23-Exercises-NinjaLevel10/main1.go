package main

import "fmt"

func main() {
	//buffer c := make(chan int, 1)

	c := make(chan int)

	go func() {
		c <- 5
	}()

	fmt.Println(<-c)
}
