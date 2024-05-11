package main

import (
	"fmt"
	"strings"
)

func checkUniq(str string) bool {
	str = strings.ToLower(str)
	checklist := make(map[rune]bool)
	for _, i := range str {
		if checklist[i] {
			return false
		}
		checklist[i] = true
	}
	return true
}

func Ex26() {
	box := "string"
	box2 := "stringS"
	fmt.Println(checkUniq(box))
	fmt.Println(checkUniq(box2))
}
