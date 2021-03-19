package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{15, 8, 1, 6}

	s := []string{"JS", "LH", "PG", "AS"}

	sort.Ints(x)
	sort.Strings(s)

	fmt.Println(x)
	fmt.Println(s)
}
