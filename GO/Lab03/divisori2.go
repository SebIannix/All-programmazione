package main

import (
	"fmt"
	)
	
	var num, div, numDiv, sumDiv int
	
func main() {
	fmt.Print("Inserire numero: ")
	fmt.Scan(&num)
	fmt.Print("Divisori di ", num, ": ")
	
	//Check divisori
	for div = 1; div < num; div++ {
		if num%div == 0 {
			fmt.Print(div, " ")
			numDiv++;
			sumDiv += div
			}
		}
		fmt.Println()
		fmt.Println("Somma dei", numDiv, "divisori:", sumDiv)
}
