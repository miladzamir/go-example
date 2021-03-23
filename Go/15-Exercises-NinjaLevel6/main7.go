package main

import "fmt"

func main() {
	x := func() string {
		return "Hello!"
	}()

	fmt.Print(x, "\n")
}
