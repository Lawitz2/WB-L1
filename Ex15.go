package main

import "fmt"

var justString string

func createHugeString(n int) string {
	var box string
	for i := 0; i < n; i++ {
		box += "большая строка"
	}
	return box
}

func someFunc() {
	v := createHugeString(1 << 10)
	length := 14
	//justString = v[:100]
	box := []rune(v)
	box = box[:min(len(box), length)]
	justString = string(box)
}

func Ex15() {
	someFunc()
	fmt.Println(justString)
}
