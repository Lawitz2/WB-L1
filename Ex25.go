package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(t int) {
	sleepchan := make(chan struct{})
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(t))
	defer ctxCancel()

	go func() {
		<-ctx.Done()
		sleepchan <- struct{}{}
		close(sleepchan)
	}()

	<-sleepchan
}

func Ex25() {
	start := time.Now()
	sleep(3)
	fmt.Printf("i've slept for %v", time.Since(start))
}
