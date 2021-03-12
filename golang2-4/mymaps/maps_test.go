package mymaps

import (
	"runtime"
	"strconv"
	"testing"
)

func BenchmarkSetMutex19(b *testing.B) {
	set := NewSetMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%10 == 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}

func BenchmarkSetMutex55(b *testing.B) {
	set := NewSetMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%2 == 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}

func BenchmarkSetMutex91(b *testing.B) {
	set := NewSetMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%10 != 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}

func BenchmarkSetRMutex19(b *testing.B) {
	set := NewSetRMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%10 == 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}

func BenchmarkSetRMutex55(b *testing.B) {
	set := NewSetRMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%2 == 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}

func BenchmarkSetRMutex91(b *testing.B) {
	set := NewSetRMutex()
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for i := 1; pb.Next(); i++ {
				if i%10 != 0 {
					set.Add(float64(i))
				} else {
					set.Has(float64(i) / 10)
				}
			}
		})
	})
}
