package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	var box string
	for i := 0; i < n; i++ {
		box += "большая строка 😁 "
	}
	return box
}

// функция, позволяющая взять часть большой строки и не держать большую строку в памяти
func someFunc() {
	v := []rune(createHugeString(1 << 10))
	length := 23
	//justString = v[:100]
	justString = strings.Clone(string(v[:min(len(v), length)]))

}

func Ex15() {
	someFunc()
	fmt.Println(justString)
}
