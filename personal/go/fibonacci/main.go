package main

import (
	"fmt"
	"math/big"
	"time"
)

// fast but inaccurate around 93
func FibonacciLoop(n uint64) uint64 {
	f := make([]uint64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	var i uint64
	for i = 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// loop style with big.Int
func BigFibonacciLoop(n int) *big.Int {
	//length := n.Add(n, 1)
	// create slice
	f := make([]*big.Int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = big.NewInt(0)
	f[1] = big.NewInt(1)
	for i := 2; i <= n; i++ {
		f[i] = big.NewInt(0)
		f[i] = f[i].Add(f[i-1], f[i-2])
	}
	return f[n]
}

// quite a bit slower
func FibonacciRecursion(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// FibonacciRecursion inaccurate at 94
func FibonacciTailRecursion(n, a, b uint64) uint64 {
	switch n {
	case 0:
		return a
	case 1:
		return b
	default:
		return FibonacciTailRecursion(n-1, b, a+b)
	}
}

// good for bigger numbers (n >= 94)
func BigFibonacciTailRecursion(n int, a, b *big.Int) *big.Int {
	switch n {
	case 0:
		return a
	case 1:
		return b
	default:
		a.Add(a, b)
		return BigFibonacciTailRecursion(n-1, b, a)
	}
}

func main() {
	fmt.Println("Loop style")
	for i := 0; i <= 100; i++ {
		start := time.Now()
		ret := FibonacciLoop(uint64(i))
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("#%d %d (%d ns)\n", i, ret, elapsed)
	}
	fmt.Println("\nBig Loop style")
	for i := 0; i <= 100; i++ {
		start := time.Now()
		ret := BigFibonacciLoop(i)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("#%d %d (%d ns)\n", i, ret, elapsed)
	}
	fmt.Println("\nTail recursion style")
	for i := 0; i <= 100; i++ {
		start := time.Now()
		ret := FibonacciTailRecursion(uint64(i), 0, 1)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("#%d %d (%d ns)\n", i, ret, elapsed)
	}
	fmt.Println("\nBig Tail recursion style")
	for i := 0; i <= 1000; i++ {
		start := time.Now()
		ret := BigFibonacciTailRecursion(i, big.NewInt(0), big.NewInt(1))
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("#%d %d (%d ns)\n", i, ret, elapsed)
	}
	fmt.Println("\nRegular recursion style")
	for i := 0; i <= 100; i++ {
		// delegate work, fill your CPUs. CAREFUL this will eat CPU
		go fmt.Printf("%d ", FibonacciRecursion(uint64(i)))
	}
}
