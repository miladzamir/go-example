package main

import "fmt"

func main() {
	c := make(chan int)

	v, ok := <-c
		fmt.Println(v, ok)

	close(c)

	v, ok = <-c
		fmt.Println(v, ok)
}
