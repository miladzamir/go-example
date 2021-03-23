package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-exist.txt")
	if err != nil {
		//fmt.Println(err)
		//log.Println(err)
		//log.Fatalln(err)
		panic(err)
	}
}
func foo() {
	fmt.Println("?")
}