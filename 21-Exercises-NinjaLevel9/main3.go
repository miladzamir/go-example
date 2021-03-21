package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	times := 10

	wg.Add(times)

	for i := 0; i < times; i++  {
		fmt.Println("A:", counter)
		v := counter
		runtime.Gosched()
		v++
		fmt.Println("B:", counter)
		counter = v
		wg.Done()
	}
	wg.Wait()

	fmt.Println("E", counter)
}