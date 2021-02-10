package main

import "fmt"

func main() {
	var num int = 42
	fmt.Printf("%d \t %b \t %#x \n", num, num, num)

	nextNum := num << 1
	fmt.Printf("%d \t %b \t %#x \n", nextNum, nextNum, nextNum)
}
