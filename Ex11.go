package main

import (
	"fmt"
	"math/rand"
)

// данная программа находит элементы, содержащиеся в обоих слайсах
// и выносит их в новый слайс
func Ex11() {
	var arr1, arr2 []int
	intersection := make(map[int]bool)

	for i := 0; i < 10; i++ {
		arr1 = append(arr1, rand.Intn(20))
		arr2 = append(arr2, rand.Intn(20))
	}
	fmt.Printf("arr 1: %v\n", arr1)
	fmt.Printf("arr 2: %v\n", arr2)

	for i := range arr1 {
		intersection[arr1[i]] = false
	}

	for i := range arr2 {
		_, ok := intersection[arr2[i]]
		if ok {
			intersection[arr2[i]] = true
		}
	}

	fmt.Printf("intersecting elements: ")
	for i := range intersection {
		if intersection[i] {
			fmt.Printf("%d ", i)
		}
	}
}
