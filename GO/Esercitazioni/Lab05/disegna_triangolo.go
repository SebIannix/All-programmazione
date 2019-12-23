package main

import (
	"fmt"
)

func main() {
	var t string = "*"
	var n int
	
	fmt.Print("Inserire intero positivo: ")
	fmt.Scan(&n)
	
	for n <= 0 {
		fmt.Print("Inserire intero positivo: ")
		fmt.Scan(&n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(t)
		}
		fmt.Println()
	}
}
