package main

import "fmt"

func main() {
	n := "Zam1r"
	switch n {
	case "Zam1r", "Abc":
		fmt.Println("name: ", n)
		fallthrough
	case "Naser":
		fmt.Println(n)

	default:
		fmt.Println("damn")
	}

}
