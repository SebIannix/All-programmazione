package main

import (
	"fmt"
	)

	var in int
	var ast string = "*"
	var più string = "+"
	
func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&in)

	for i := 0; i < in; i++ {	
		for j := 1; j <= in; j++ {
			if j%2 == 0 {
				fmt.Print(più)
			} else {
				fmt.Print(ast)
			}
		}
		
		fmt.Println()
	}
}
