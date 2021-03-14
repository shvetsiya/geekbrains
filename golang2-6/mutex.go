package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
