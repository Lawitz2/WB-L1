package main

import "fmt"

// данная программа группирует слайс температур в группы с шагом в 10 градусов, где 10=[10-20)
func Ex10() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempmap := make(map[int][]float64)

	for i := range temps {
		switch {
		case temps[i] < 0:
			tempmap[int((temps[i]-10)/10)*10] = append(tempmap[int((temps[i]-10)/10)*10], temps[i])
		default:
			tempmap[int(temps[i]/10)*10] = append(tempmap[int(temps[i]/10)*10], temps[i])

		}
	}
	fmt.Println(tempmap)
}
