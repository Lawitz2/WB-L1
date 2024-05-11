package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counterEx18 struct {
	atomic.Uint64
}

func worker18(wg *sync.WaitGroup, jobsAmount *counterEx18) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//job
		jobsAmount.Add(1)
	}
}

func Ex18() {
	jobsAmount := counterEx18{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker18(wg, &jobsAmount)
	}
	wg.Wait()
	fmt.Printf("we did a total of %v jobs", jobsAmount.Load())
}
