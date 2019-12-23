package main

import (
	"fmt"
)

func main() {
	var r string = "*"
	var a, b int
	
	fmt.Print("Inserire base del rettangolo: ")
	fmt.Scan(&b)
	fmt.Print("Inserire altezza del rettangolo: ")
	fmt.Scan(&a)
	
	for a <= 0 || b <= 0 {
		fmt.Print("Inserire intero positivo come base del rettangolo: ")
		fmt.Scan(&b)
		fmt.Print("Inserire intero positivo come altezza del rettangolo: ")
		fmt.Scan(&a)
	}
	
	for j := 0; j < a; j++ {
		for i := 0; i < b; i++ {
			fmt.Print(r)
		}
		fmt.Println()
	}
}
