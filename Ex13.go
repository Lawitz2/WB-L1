package main

import "fmt"

func Ex13() {
	a := 5
	b := 10
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b)
}
