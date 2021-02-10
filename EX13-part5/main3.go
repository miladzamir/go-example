package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	tr := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}

	se := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(tr)
	fmt.Println(se)

	fmt.Println(tr.doors, tr.color, tr.fourWheel)

	fmt.Println(se.doors, se.color, se.luxury)
}
