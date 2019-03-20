package main

import (
	"fmt"
	"math"
	)

	var in, valori, tot int

func main() {
	fmt.Print("Inserire quantità di numeri: ")
	fmt.Scan(&in)
	
	var list = make([]float64, in)
	
	fmt.Print("Inserire valori strettamente positivi: ")
	for valori = range list {
		fmt.Scan(&valori)
	
		if valori > 0 {
			tot += valori
		}
	}
	
	var x float64 = float64(tot)
	var x2 float64 = math.Pow(x, 2)
	var n float64 = float64(in)
	
	//fmt.Println("Somma:", tot)
	fmt.Println()
	fmt.Println("	Media Aritmetica:", x/n)
	fmt.Println("	Media Geometrica:", math.Sqrt(x))
	fmt.Println("	Media Quadratica:", math.Sqrt(x2/n))
	fmt.Println("	Media Armonica:", n/(1/x))
}
