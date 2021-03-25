package main

import (
	"fmt"

	"github.com/shvetsiya/geekbrains/golang1-04/fibonacci"
)

func main() {
	memo := make(map[int64]int64)

	var n int64 = 1
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci.Fibonacci0(n), fibonacci.Fibonacci0(n) == fibonacci.Fibonacci1(n, memo))

	n = 2
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci.Fibonacci0(n), fibonacci.Fibonacci0(n) == fibonacci.Fibonacci1(n, memo))

	n = 20
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci.Fibonacci0(n), fibonacci.Fibonacci0(n) == fibonacci.Fibonacci1(n, memo))
}
