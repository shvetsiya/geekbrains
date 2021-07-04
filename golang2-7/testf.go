package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shvetsiya/geekbrains/golang2-4/mymaps"
)

func main() {
	atomicIncrement()
	mutexIncrement()

	setMutex := mymaps.NewSetMutex()
	//setRMutex := mymaps.NewSetRMutex()
	stime := time.Now()
	for i := 0; i < 1000; i++ {
		setMutex.Add(float64(i))
	}
	fmt.Println(stime)
}

func atomicIncrement() {
	var wg sync.WaitGroup
	sum := int32(0)
	wg.Add(1)

	go func() {
		atomic.AddInt32(&sum, 1)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(sum)
}

func mutexIncrement() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	sum := 0
	wg.Add(1)
	go func() {
		mutex.Lock()
		sum++
		mutex.Unlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(sum)
}
