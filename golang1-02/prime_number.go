package main

import "fmt"

func main() {
	// We use Eratosthenes sieve to find prime numbers
	// Complexity: Time - O(n), Space - O(n)
	n := 30 // find all primes in the range 2..n
	primes := findPrimes(n)
	for i := 2; i < len(primes); i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}

}

func findPrimes(n int) []bool {
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {

			for j := i; j*i < n; j++ {
				primes[i*j] = false
			}
		}
	}
	return primes
}
