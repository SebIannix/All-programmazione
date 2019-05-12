package main

import (
	"fmt"
	)

	var in int
	var ast string = "*"
	var più string = "+"
	var counter int
	
func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&in)
	
	for i := 0; i < in; i++ {
		//fmt.Println("DEBUG: 'for' interno, caratteri")
		for j := 1; j <= in; j++ {
			
		}
		
		fmt.Println()
	}
}

			/* //dopo 'for' interno
			if i%2 == 0 {
				//fmt.Println("DEBUG: righe pari")
				for counter = 0; counter < 4; counter++ {
					if counter <= 1 {
						fmt.Print(più)
					} else if counter > 1 {
						fmt.Print(ast)
						if counter == 4 {
							counter = 0
						}
					}
				}
			} else {
				//fmt.Println("DEBUG: righe dispari")
				for counter = 0; counter < 4; counter++ {
					if counter <= 1 {
						fmt.Print(ast)
					} else if counter > 1 {
						fmt.Print(più)
						if counter == 4 {
							counter = 0
						}
					}
				}
			}
			*/
