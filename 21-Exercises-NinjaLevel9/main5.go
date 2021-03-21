package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64
	times := 10

	wg.Add(times)

	for i := 0; i < times; i++  {
		atomic.AddInt64(&counter, 1)
		fmt.Println("B:", atomic.LoadInt64(&counter))
		wg.Done()
	}
	wg.Wait()

	fmt.Println("E", counter)
}