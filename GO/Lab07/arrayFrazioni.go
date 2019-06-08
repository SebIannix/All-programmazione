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
	fmt.Print("Output: ", x, "\n")
	
	for i := 0; i < len(x); i++ {
		x[i] = 
	}
}
