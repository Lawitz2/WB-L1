package main

import (
	"fmt"
	"sync"
)

func half(wg *sync.WaitGroup, arr *[]int, in chan int) { // я отправляю данные из слайса в канал 1
	defer wg.Done()
	for i := range *arr {
		in <- (*arr)[i]
	}
	close(in)
}

func otherHalf(wg *sync.WaitGroup, in, out chan int) { // я получаю данные из канала 1, возвожу в квадрат и отправляю в канал 2
	defer wg.Done()
	for box := range in {
		out <- box * box
	}
	close(out)
}

func printout(wg *sync.WaitGroup, out chan int) { // я получаю данные из канала 2 и вывожу их
	defer wg.Done()
	for box := range out {
		fmt.Printf("%d ", box)
	}
}

func Ex9() {
	wg := &sync.WaitGroup{}

	inChan := make(chan int, 1)
	outChan := make(chan int, 1)

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(3)

	go half(wg, &arr, inChan)
	go otherHalf(wg, inChan, outChan)
	go printout(wg, outChan)

	wg.Wait()
}
