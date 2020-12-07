package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		res := ""
		if i%3 == 0 {
			res += "Fizz"
		}
		if i%5 == 0 {
			res += "Buzz"
		}
		if res != "" {
			fmt.Println(res)
			continue
		}
		fmt.Println(i)
	}
}
