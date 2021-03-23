package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)
}
func receive(c, q <-chan int)  {
	for {
		select {
		case v := <-c:
			fmt.Println("C:", v)
		case v := <-q:
			fmt.Println("Q:", v)
			return
		}
	}
}
func gen(q chan <- int) <-chan int {
	c := make(chan int)

	go func() {
		for i:=0; i<100; i++ {
			c <-i
		}
		q <- 1
		close(c)
	}()
	return c
}
