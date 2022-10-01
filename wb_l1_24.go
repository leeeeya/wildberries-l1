// Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// конструктор
func (p *Point) setPoint(x, y float64) {
	p.y, p.x = y, x
}

// вычисление расстояния между двумя точками
func (p Point) getDistance() float64 {
	return math.Abs(p.y - p.x)
}

func main() {
	p := new(Point)
	p.setPoint(3.45, -5.678)
	fmt.Printf("Distace between x and y: %.2f\n", p.getDistance())
}
