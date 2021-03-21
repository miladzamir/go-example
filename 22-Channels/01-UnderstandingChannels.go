package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 50
	}()

	fmt.Println(<-c)

	//Buffer , 1
	c2 := make(chan int, 2)

	c2 <- 41
	c2 <- 51

	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
