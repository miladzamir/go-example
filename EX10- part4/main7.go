package main

import "fmt"

func main() {
	//create slice of slice ??!! :/

	x := [][]string{}

	y := []string{"james", "bound", "shaken"}
	z := []string{"miss", "moneyPeeny", "Hello!!!"}

	// another Way
	// x := [][]string{x, y}

	x = append(x, y, z)
	fmt.Println(x)

	fmt.Println("-/-/-/-/-/-/-/-/-/")
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Println("-/-/-/-/-/-/-/-/-/")
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Println(x[i][j])
		}
	}
}
