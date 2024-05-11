package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
	Способ завершения горутин через закрытие канала с работами был выбран потому,
	что в данном случае нужно писать меньше всего кода, и меньше шансов ошибиться.
*/

func worker(w int, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("i am worker18 #%d, i got \"%v\" from the main channel\n", w, job)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
	}
	fmt.Printf("i am worker18 #%d, i am cleaning up and stopping\n", w)
}

func Ex4() {
	continueFeedingJobs := true
	wg := &sync.WaitGroup{}

	go func() {
		/*
			Рутина смотрит за интеррапт сигналом (ктрл+С), когда он поступает - перестаём
			создавать работы и закрываем канал, что заставит воркеров по окончании всех работ
			выйти из цикла и завершиться
		*/
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, os.Interrupt)
		<-termChan
		fmt.Println("termination invoked")
		continueFeedingJobs = false
		close(termChan)
	}()

	var numOfWorkers int
	fmt.Printf("num of workers: ")
	_, err := fmt.Scanf("%d", &numOfWorkers)
	if err != nil {
		slog.Error("couldn't read from stdin: ", err.Error())
		return
	}

	if numOfWorkers <= 0 {
		fmt.Println("incorrect number of workers")
		return
	}

	workerFeeder := make(chan any, 1)

	wg.Add(numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		go worker(i, workerFeeder, wg)
	}

	for continueFeedingJobs {
		workerFeeder <- rand.Intn(500)
		//time.Sleep(time.Millisecond * time.Duration(500+rand.Intn(1000)))
	}
	close(workerFeeder)

	wg.Wait()
	fmt.Println("i am main thread, all workers are terminated, shutting down.")
}
