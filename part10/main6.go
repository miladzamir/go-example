package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(x)

	//args x, ...T
	x = append(x, 11, 22, 33, 44, 55, 66, 77, 88, 99)
	fmt.Println(x)

	y := []int{111, 222, 333, 444, 555, 666, 777, 888, 999}

	x = append(x, y...)
	fmt.Println(x)

	//delete slice
	//REMOVE 20 30 from x slice
	x = append(x[:1], x[3:]...)
	fmt.Println(x)

}
