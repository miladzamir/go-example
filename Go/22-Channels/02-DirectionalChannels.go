package main

import "fmt"

func main() {
	cs := make(chan<- int) //send
	cr := make(<-chan int) // receive

	fmt.Printf("%T \n", cs)
	fmt.Printf("%T \n", cr)

}
