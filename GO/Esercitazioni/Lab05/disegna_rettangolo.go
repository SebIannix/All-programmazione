package main

import (
	"fmt"
)

func main() {
	var r string = "*"
	var a, b int
	
	fmt.Print("Inserire base rettangolo: ")
	fmt.Scan(&b)
	fmt.Print("Inserire altezza rettangolo: ")
	fmt.Scan(&a)
	
	for j := 0; j < a; j++ {
		for i := 0; i < b; i++ {
			fmt.Print(r)
		}
		fmt.Println()
	}
}
