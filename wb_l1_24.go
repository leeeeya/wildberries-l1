package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) setPoint(x, y float64) {
	p.y, p.x = y, x
}

func (p Point) getDistance() float64 {
	return math.Abs(p.y - p.x)
}

func main() {
	p := new(Point)
	p.setPoint(3.45, -5.678)
	fmt.Printf("Distace between x and y: %.2f\n", p.getDistance())
}
