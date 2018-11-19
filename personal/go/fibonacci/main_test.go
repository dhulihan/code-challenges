package main

import (
	"fmt"
	"testing"
)

func benchmarkFibonacciTailRecursion(i int, b *testing.B) {
	fmt.Println("Running", b.N, "times")
	for n := 0; i <= b.N; n++ {
		FibonacciTailRecursion(uint64(i), 0, 1)
	}
}

func BenchmarkFibonacciTailRecursion(b *testing.B) { benchmarkFibonacciTailRecursion(10, b) }
