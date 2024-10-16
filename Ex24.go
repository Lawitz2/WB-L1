package main

import (
	"fmt"
	"math"
)

type point struct {
	X float64
	Y float64
}

func newPoint(x float64, y float64) *point {
	return &point{X: x, Y: y}
}

func distance(a *point, b *point) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + (math.Pow(b.Y-a.Y, 2)))
}

// вычислени расстояния между двумя точками на 2д-пространстве
func Ex24() {
	a := newPoint(2, 2)
	b := newPoint(0, 0)

	fmt.Println(distance(a, b))
}
