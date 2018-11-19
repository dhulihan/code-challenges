package main

import (
	"fmt"
	"strconv"
)

// fast but inaccurate around 93
func FibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// quite a bit slower
func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Print(strconv.Itoa(FibonacciLoop(i)) + " ")
	}
	fmt.Println("")
	for i := 0; i <= 100; i++ {
		// delegate work, fill your CPUs. CAREFUL this will eat CPU
		go fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
	}
	fmt.Println("")
}
