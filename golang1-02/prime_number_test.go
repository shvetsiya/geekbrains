package main

import (
	"reflect"

	"testing"
)

func TestFindPrimes(t *testing.T) {
	primes := findPrimes(4)
	test := []bool{false, false, true, true}
	if !reflect.DeepEqual(primes, test) {
		t.Errorf("bsdfsdf")
	}
}

func TestFindPrimesLen(t *testing.T) {
	primes := findPrimes(10)
	if len(primes) != 10 {
		t.Errorf("The size of the array is incorrect")
	}
}
