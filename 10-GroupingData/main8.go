package main

import "fmt"

func main() {
	pr1 := []string{"ali", "mohammdi", "19", "iran"}
	fmt.Println(pr1)
	pr2 := []string{"reza", "Naseri", "23", "afghanestan"}
	fmt.Println(pr2)

	people := [][]string{pr1, pr2}
	fmt.Println(people)
	fmt.Println("")
	fmt.Println(people[0])
}
