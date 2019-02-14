package main

import (
	"fmt"
	"math"
)

var n, ints, sum int
var over, under, z  = 0, 0, 0
var min = math.MinInt64
var max = math.MaxInt64

func main() {
	fmt.Print("Inserisci quantità di interi: ")
	fmt.Scan(&n)

	var a = make([]int, n)

	fmt.Println("Inserisci", n, "numeri:")
	for ints = range a {
		fmt.Scan(&ints)
		
		
		//Somma
		sum += ints
		
		//Valore massimo
		if ints > min {
			min = ints
		}
		
		//Valore minimo
		if ints < max {
			max = ints
		}
		
		//Valori
		
	}
	fmt.Println("Somma:", sum)
	fmt.Println("Valore minimo:", max)
	fmt.Println("Valore massimo:", min)
	fmt.Println("Interi > 0:", over)
	fmt.Println("Interi < 0:", under)
	fmt.Println("Interi = 0:", z)

	fmt.Println("DEBUG:", a)
}
