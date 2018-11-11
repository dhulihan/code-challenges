package main

import "fmt"

func FirstFactorial(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum

}

// Enable for testing
//func readline() int {
//return 0
//}

func main() {
	fmt.Println(FirstFactorial(readline()))
}
