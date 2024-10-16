package main

import (
	"fmt"
	"math"
	"math/big"
)

// работа с большими числами (вне пределов int64)
func Ex22() {
	base := 2.0
	power := 62.0
	a := big.NewInt(int64(math.Pow(base, power)))
	b := big.NewInt(int64(math.Pow(base, power-1)))
	mul := big.NewInt(0)
	mul.Mul(a, b)
	div := big.NewInt(0)
	div.Div(a, b)
	add := big.NewInt(0)
	add.Add(a, b)
	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Printf("a: %d, b: %d\na*b: %d\na/b: %d\na+b: %d\na-b: %d", a, b, mul, div, add, sub)
}
