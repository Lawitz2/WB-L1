package main

import (
	"fmt"
	"sync"
)

func squares(wg *sync.WaitGroup, s int) {
	defer wg.Done()
	fmt.Printf("%d ", s*s)
}

func Ex2() {
	wg := sync.WaitGroup{}

	numbers := []int{2, 4, 6, 8, 10}

	wg.Add(len(numbers)) // добавляем ВГ по кол-ву го рутин

	for i := 0; i < len(numbers); i++ {
		go squares(&wg, numbers[i])
	}

	wg.Wait() // ждём завершения всех го рутин
}
