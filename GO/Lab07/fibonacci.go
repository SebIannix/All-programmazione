package main

import(
	"fmt"
)

var x int

func Fibonacci(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	
	var n2, n1 uint64 = 0, 1
	
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
		fmt.Print(n1, " ")
	}
	return n2 + n1
}

func main() {
	Fibonacci(51)
	fmt.Println()
}
