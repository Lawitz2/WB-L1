package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

func workerEx7(ctx context.Context, lock *sync.Mutex, index int, counter map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			lock.Lock()
			counter[index]++
			lock.Unlock()
		}
	}
}

// реализация конкурентного доступа к мапе, завершающаяся через заданное время
func Ex7() {
	counter := make(map[int]int)
	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	var seconds int

	fmt.Printf("seconds to work: ")
	_, err := fmt.Scanf("%d", &seconds)
	if err != nil {
		slog.Error("couldn't read from stdin: ", err.Error())
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancelFunc()

	start := time.Now()

	index := 1
	workerPoolSize := 2
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go workerEx7(ctx, lock, index, counter, wg)
	}

	wg.Wait()
	fmt.Printf("in %v we increased index %d in a map %d times", time.Since(start), index, counter[index])
}
