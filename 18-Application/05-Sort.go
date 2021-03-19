package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{5, 9, 5, 1, 6, 9, 51, 99, 15, 6}

	s := []string{"d", "h", "e", "f", "p", "j", "a", "b"}

	sort.Ints(x)
	sort.Strings(s)

	fmt.Println(x)
	fmt.Println(s)
}
