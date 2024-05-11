package main

import (
	"fmt"
	"slices"
	"strings"
)

func wordOrder(s string) string {
	box := strings.Split(s, " ")
	slices.Reverse(box)
	return strings.Join(box, " ")
}

func Ex20() {
	box := "this is time ⌛ and this is a smile 😁 and this is немного кириллицы"
	box = wordOrder(box)
	fmt.Println(box)
}
