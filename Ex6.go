package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func routine1(stopChannel chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChannel:
			fmt.Println("i am routine 1, i am stopping because of signal from a channel")
			return
		default:
			// jobs
			time.Sleep(time.Second)
		}
	}
}

func routine2(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for !*stop {
		// jobs
		time.Sleep(time.Second)
	}
	fmt.Println("i am routine 2, i am stopping because someone swapped my bool flag")
	return
}

func routine3(jobs chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for range jobs {
		// jobs
		time.Sleep(time.Second)
	}
	fmt.Println("i am routine 3, i am stopping because someone closed the channel that i was getting my jobs from")
	return
}

func routine4(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("i am routine 4, i am stopping because someone invoked context cancel func. It can be manual call, or i can timeout.")
			return
		default:
			//jobs
			time.Sleep(time.Second)
		}
	}
}

func routine5(wg *sync.WaitGroup) {
	defer wg.Done()
	//jobs
	fmt.Println("i am routine 5, i am stopping because i don't have any blocks/loops and i simply went through all of my code lines.")
	return
}

// пример разных способов завершения горутин
func Ex6() {
	stopSignalChannel := make(chan struct{})                   //1
	stop := false                                              //2
	workerFeeder := make(chan any)                             //3
	ctx, cancelCtx := context.WithCancel(context.Background()) //4

	wg := &sync.WaitGroup{}
	wg.Add(5)

	go routine1(stopSignalChannel, wg)

	go routine2(&stop, wg)

	go routine3(workerFeeder, wg)

	go routine4(ctx, wg)

	go routine5(wg)

	stopSignalChannel <- struct{}{}
	stop = true
	close(workerFeeder)
	cancelCtx()

	wg.Wait()
	fmt.Println("all routines are terminated")
}
