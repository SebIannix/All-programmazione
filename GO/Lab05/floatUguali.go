package main

import (
	"fmt"
	"math"
)

var x, y float64

func main() {
	fmt.Print("Inserire valore: ")
	fmt.Scan(&x)
	
	y = math.Sqrt(x)
	if y*y == x {
		fmt.Print(x, " è uguale a ", y, "*", y, "\n")
	} else {
		fmt.Print(x, " è diverso da ", y, "*", y, "\n")
	}
}
