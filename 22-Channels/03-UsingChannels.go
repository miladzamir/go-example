package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)
}

func foo(c chan<- int) {
	c <- 51
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
