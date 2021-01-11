package main

import "fmt"

func main() {
	memo := make(map[int64]int64)

	var n int64 = 1
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci0(n), fibonacci0(n) == fibonacci1(n, memo))

	n = 2
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci0(n), fibonacci0(n) == fibonacci1(n, memo))

	n = 20
	fmt.Printf("Fibonacci[%d] = %d and (Fib0 == Fib1) is %t\n", n, fibonacci0(n), fibonacci0(n) == fibonacci1(n, memo))
}

// fibonacci0 finds nth Fibonacci number
// It uses recursive formula to find the numbers: Fn = F(n-1) + F(n-2)
//
// Special cases are:
// F(0) = 0
// F(1) = 1
// for more details see https://en.wikipedia.org/wiki/Fibonacci_number
func fibonacci0(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci0(n-1) + fibonacci0(n-2)
}

// fibonacci1 finds nth Fibonacci number
// It uses recursive formula to find the numbers: Fn = F(n-1) + F(n-2)
// as in fibonacci0, but uses optimization based on memoization i.e
// if we already saw this number before immediately return it without
// doing any calculations
//
// Special cases are:
// F(0) = 0
// F(1) = 1
// if n in memo: return memo[n]
// for more details see https://en.wikipedia.org/wiki/Fibonacci_number
func fibonacci1(n int64, memo map[int64]int64) int64 {
	v, isExit := memo[n]
	if isExit == true {
		return v
	}

	if n == 0 || n == 1 {
		return n
	}
	memo[n] = fibonacci1(n-1, memo) + fibonacci1(n-2, memo)
	return memo[n]
}
