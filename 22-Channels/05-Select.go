package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("EVEN:", v)
		case v := <-o:
			fmt.Println("ODD:", v)
		case v := <-q:
			fmt.Println("Quit:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}
