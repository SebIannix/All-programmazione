package main

import (
	"fmt"
	
)

const dim = 5

func switchFirstLast(a [dim]int) [dim]int {
	for i, j := 1, len(a)-2; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func reverse(a [dim]int) [dim]int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	array := [dim]int{1, 2, 3, 4, 5}
	
	fmt.Println("Array:", array)
	fmt.Println("Reverse array:", reverse(array))
	fmt.Println("Switch First-Last:", switchFirstLast(array))
}
