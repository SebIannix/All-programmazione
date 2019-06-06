package main

import (
	"fmt"
)

var a, b float64

func main() {
	fmt.Print("Equazione generica: a*x + b = 0 \nInserire valore a: ")
	fmt.Scan(&a)
	fmt.Print("Inserire valore b: ")
	fmt.Scan(&b)
	
	x := -b/a
	fmt.Println("Risultato: ", x)
}
