package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 5
		//close(c)
	}()

	i, ok := <-c
	fmt.Println(i, ok)

	//i, ok = <-c
	//fmt.Println(i, ok)
}
