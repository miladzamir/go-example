package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {

	p1 := person{
		first: "Sara",
		last:  "Ja",
		favoriteIceCream: []string{
			"kramel",
			"new",
			"basaniiii",
		},
	}
	p2 := person{
		first: "james",
		last:  "bound",
		favoriteIceCream: []string{
			"ja",
			"fav",
			"BAAA",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first, v.last)
		for _, v2 := range v.favoriteIceCream {
			fmt.Println("\t", v2)
		}
	}

	fmt.Println("")

	fmt.Println("***=>", p1.first, p1.last)
	for _, v := range p1.favoriteIceCream {
		fmt.Println(v)
	}

	fmt.Println("***=>", p2.first, p2.last)

	for _, v := range p2.favoriteIceCream {
		fmt.Println(v)
	}
}
