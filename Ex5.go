package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"
)

func Ex5() {
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

	ch := make(chan int, 1)

loop:
	for {
		ch <- rand.Intn(1000)
		select {
		case job := <-ch:
			fmt.Printf("got %d from channel\n", job)
		case <-ctx.Done():
			close(ch)
			break loop
		}
		time.Sleep(20 * time.Millisecond)
	}

	fmt.Printf("we're done, i spent %v doing jobs\n", time.Since(start))
}
