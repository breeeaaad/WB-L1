package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func New(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) distance(o Point) float64 {
	//формула нахождения расстояния
	return math.Sqrt(math.Pow(float64(p.x-o.x), 2) + math.Pow(float64(p.y-o.y), 2))
}

func main() {
	p1 := New(1, 5)
	p2 := New(4, 6)
	fmt.Print(p1.distance(*p2))
}
