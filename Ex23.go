package main

import "fmt"

func Ex23() {
	a := []int{1, 2, 3, 4, 5} // если важно сохранить порядок
	delPosA := 2
	a = append(a[:delPosA], a[delPosA+1:]...)
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5} // если порядок не важен
	delPosB := 1
	b[delPosB], b[len(b)-1] = b[len(b)-1], b[delPosB]
	b = b[:len(b)-1]
	fmt.Println(b)
}
