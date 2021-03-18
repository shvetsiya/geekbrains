package main

import "fmt"

func main() {
	var counters = map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, k int) {
			for j := 0; j < 5; j++ {
				counters[k*10+j]++
			}
		}(counters, i)
	}
	fmt.Println("counters result", counters)
}
