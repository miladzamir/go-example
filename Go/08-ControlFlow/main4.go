package main

import "fmt"

func main() {
	//single condition

	a := 1
	b := 5

	for a <= b {
		fmt.Printf("a is : %v b is :%v \n", a, b)
		a++
	}

	fmt.Println("-/-/-/-/-/-/-/-/-/-")

	//for statements with for cause

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-/-/-/-/-/-/-/-/-/-")

	//for like while true
	count := 1
	for {
		if count < 10 {
			count++
			fmt.Println(count)
		} else {
			break
		}
	}
}
