package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Retangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Retangle) area() float64 {
	return r.width * r.height
}
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Retangle{width: 15, height: 20}

	fmt.Printf("circle area: %f\n", getArea(circle))
	fmt.Printf("rectangle area: %f\n", getArea(rectangle))
}
