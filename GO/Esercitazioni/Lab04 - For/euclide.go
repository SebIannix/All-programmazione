package main

import (
	"fmt"
)

var a, b, r int
var i int = 1

func main(){
	fmt.Print("Inserire interi naturali: ")
	fmt.Scan(&a, &b)	//input valori
	
	for a < 0 || b < 0{	//check valori naturali
		fmt.Print("Inserire interi naturali: ")
		fmt.Scan(&a, &b)
	}
	
	for b != 0 {	//algoritmo Euclide
		fmt.Printf("%v) - Resto di %v/%v: %v\n", i, a, b, r)
		r = a % b
		a = b
		b = r
		i++
	}
	fmt.Println("MCD:", a)
}
