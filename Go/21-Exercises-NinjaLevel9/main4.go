package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	times := 10

	wg.Add(times)

	var m sync.Mutex

	for i := 0; i < times; i++  {
		m.Lock()
		v := counter
		v++
		counter = v
		fmt.Println("B:", counter)
		m.Unlock()
		wg.Done()
	}
	wg.Wait()

	fmt.Println("E", counter)
}