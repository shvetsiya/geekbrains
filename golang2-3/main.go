package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const goroutinesNum = 1000

func main() {
	workerInput := make(chan struct{}, 10)
	var chanWg sync.WaitGroup
	sum := 0
	chanWg.Add(1)
	go func() {
		for range workerInput {
			sum++
		}
		chanWg.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(goroutinesNum)
	for i := 0; i < goroutinesNum; i++ {
		go func(job int) {
			workerInput <- struct{}{}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(workerInput)
	chanWg.Wait()

	fmt.Printf("Final sum is: %d\n", sum)
}

func contextWithSigterm(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithTimeout(ctx, time.Second)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

		select {
		case <-signalCh:
		case <-ctx.Done():
		}
	}()

	return ctxWithCancel
}
