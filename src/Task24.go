package main

import (
	"fmt"
	"math"
)

/* Задание: Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором. */

type Point struct {
	x float64
	y float64
}

func (self *Point) GetX() float64 {
	return self.x
}

func (self *Point) GetY() float64 {
	return self.y
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(first, second Point) float64 {
	sqr1 := math.Pow(second.GetX()-first.GetX(), 2)
	sqr2 := math.Pow(second.GetY()-first.GetY(), 2)
	return math.Sqrt(sqr1 + sqr2)
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(2, 2)
	fmt.Println(Distance(point1, point2))
}
