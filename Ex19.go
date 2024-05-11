package main

import (
	"fmt"
	"slices"
)

func reverse(s string) string {
	box := []rune(s)
	slices.Reverse(box)
	return string(box)
}

func Ex19() {
	box := "this is time ⌛, this is a smile 😁. Немного кириллицы."
	box = reverse(box)
	fmt.Println(box)
}
