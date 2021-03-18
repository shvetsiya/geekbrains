package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}
