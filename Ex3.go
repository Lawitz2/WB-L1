package main

import (
	"fmt"
	"sync"
)

func squaresSum(wg *sync.WaitGroup, i int, ch chan<- int) {
	defer wg.Done()
	ch <- i * i // конкурентно считаем и пишем в канал
}

func Ex3() {
	wg := &sync.WaitGroup{}
	var sum int

	numbers := []int{2, 4, 6, 8, 10}

	ch := make(chan int, len(numbers))

	wg.Add(len(numbers))

	for i := 0; i < len(numbers); i++ {
		go squaresSum(wg, numbers[i], ch)
	}

	for i := 0; i < len(numbers); i++ {
		sum += <-ch // суммируем всё, что приходит из канала
	}

	wg.Wait()
	close(ch)

	fmt.Println(sum)
}
