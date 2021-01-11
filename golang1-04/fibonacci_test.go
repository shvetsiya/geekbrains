package main

import (
	"testing"
)

func TestFibonacci2(t *testing.T) {
	memo := make(map[int64]int64)

	var n int64 = 2
	if fibonacci0(n) != 1 && fibonacci1(n, memo) != 1 {
		t.Errorf("fibonacci0 is not equal to fibonacci1 at n = %d", n)
	}
}

func TestFibonacci10(t *testing.T) {
	memo := make(map[int64]int64)

	var n int64 = 10
	if fibonacci0(n) != 55 && fibonacci1(n, memo) != 55 {
		t.Errorf("fibonacci0 is not equal to fibonacci1 at n = %d", n)
	}
}

func TestFibonacci15(t *testing.T) {
	memo := make(map[int64]int64)

	var n int64 = 20
	if fibonacci0(n) != fibonacci1(n, memo) {
		t.Errorf("fibonacci0 is not equal to fibonacci1 at n = %d", n)
	}
}

func BenchmarkFibonacci0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x55 := fibonacci0(10); x55 != 55 {
			b.Fatalf("Unexpected value : %d", x55)
		}
	}
}

func BenchmarkFibonacci1(b *testing.B) {
	memo := make(map[int64]int64)
	for i := 0; i < b.N; i++ {
		if x55 := fibonacci1(10, memo); x55 != 55 {
			b.Fatalf("Unexpected value : %d", x55)
		}
	}
}
