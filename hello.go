package main

import (
	"fmt"
	"math"
)

type rect struct {
	height int
	width  int
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
}

func (r *rect) area() float64 {
	return float64(r.height) * float64(r.width)
}

func (c *circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) float64 {
	return g.area()
}

func main() {
	rect_one := rect{height: 100, width: 100}

	circle_one := circle{radius: 10.5}

	fmt.Println(measure(&rect_one))
	fmt.Println(measure(&circle_one))
}
