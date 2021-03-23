package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type ByAge []Person

type ByFirst []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {

	p1 := Person{First: "a", Last: "B", Age: 25}
	p2 := Person{First: "wq", Last: "qt", Age: 34}
	p3 := Person{First: "r", Last: "t", Age: 2}
	p4 := Person{First: "aq", Last: "r", Age: 12}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByFirst(people))
	fmt.Println(people)
}
