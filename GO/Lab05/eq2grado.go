package main

import (
	"fmt"
	"math"
)

var a, b, c float64

func main() {
	
	fmt.Print("Equazione generica: a*(x^2) + b*x + c = 0\nInserire valore a: ")
	fmt.Scan(&a)
	fmt.Print("Inserire valore b: ")
	fmt.Scan(&b)
	fmt.Print("Inserire valore c: ")
	fmt.Scan(&c)
	fmt.Println()
	
	delta := b*b - 4*a*c
	if delta >= 0  {
		x1 := (-b + math.Sqrt(delta))/2*a
		x2 := (-b - math.Sqrt(delta))/2*a
		fmt.Print("Delta: ", delta, "\nPrimo risultato: ", x1, "\nSecondo risultato: ", x2, "\n")
	} else {
		
	}
}
