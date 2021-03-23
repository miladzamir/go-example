package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}
type square struct {
	length float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float32 {
	return s.length * s.length
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	ci := circle{
		radius: 12.3456,
	}

	sq := square{
		length: 15,
	}

	info(ci)
	info(sq)
}
