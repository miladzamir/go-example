package main

import "fmt"

func main() {
	var favSport string = "result"

	switch favSport {
	case "value":
		fmt.Println("value")
	case "result":
		fmt.Println("result")
	default:
		fmt.Println("!")
	}
}
