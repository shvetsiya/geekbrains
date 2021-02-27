package fibonacci

// Fibonacci0 finds nth Fibonacci number
// It uses recursive formula to find the numbers: Fn = F(n-1) + F(n-2)
//
// Special cases are:
// F(0) = 0
// F(1) = 1
// for more details see https://en.wikipedia.org/wiki/Fibonacci_number
func Fibonacci0(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci0(n-1) + Fibonacci0(n-2)
}

// Fibonacci1 finds nth Fibonacci number
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
func Fibonacci1(n int64, memo map[int64]int64) int64 {
	v, isExit := memo[n]
	if isExit == true {
		return v
	}

	if n == 0 || n == 1 {
		return n
	}
	memo[n] = Fibonacci1(n-1, memo) + Fibonacci1(n-2, memo)
	return memo[n]
}
