package main

import (
	"fmt"
	)
	
	var in1, in2, div, numDiv1, sumDiv1, numDiv2, sumDiv2 int
	
func main() {
	fmt.Print("Inserire primo numero: ")
	fmt.Scan(&in1)
	fmt.Print("Inserire secondo numero: ")
	fmt.Scan(&in2)
	
	//Check divisori primo numero
	fmt.Println()
	fmt.Print("Divisori di ", in1, ": ")
	for div = 1; div < in1; div++ {
		if in1%div == 0 {
			fmt.Print(div, " ")
			numDiv1++;
			sumDiv1 += div
			}
		}
		fmt.Println()
		fmt.Println("Somma dei", numDiv1, "divisori del primo numero:", sumDiv1)
		
	if sumDiv1 != in2 {
		fmt.Println()
		fmt.Println(in1, "e", in2, "non sono numeri amichevoli")
	} else {
		//Check divisori secondo numero
		fmt.Println()
		fmt.Print("Divisori di ", in2, ": ")
		for div = 1; div < in2; div++ {
			if in2%div == 0 {
				fmt.Print(div, " ")
				numDiv2++;
				sumDiv2 += div
				}
			}
			fmt.Println()
			fmt.Println("Somma dei", numDiv2, "divisori del secondo numero:", sumDiv2)
	}
}
