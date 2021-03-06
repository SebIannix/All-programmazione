package main

import (
	"fmt"
)

var x int

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x,d)
	}
	return input(x, err)
}

func main() {
	fmt.Print("Inserire valori: ")
	x := input([]int{}, nil)
	fmt.Print("Output:		", x, "\n")
	
	for i := len(x)/2-1; i >= 0; i-- {
		y := len(x)-1-i
		x[i], x[y] = x[y], x[i]
	}
	fmt.Print("Reverse output:	", x, "\n")
}
