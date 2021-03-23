package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(x)

	//args x, ...T
	x = append(x, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200)
	fmt.Println(x)

	y := []int{210, 220, 230, 240, 250, 260, 270, 280, 290, 300}

	x = append(x, y...)
	fmt.Println(x)

	//delete slice
	//REMOVE 20 30 from x slice
	x = append(x[:1], x[3:]...)
	fmt.Println("")
	fmt.Println(x)

}
