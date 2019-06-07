package main

import (
	"fmt"
)

var x int
var y uint64 = 1

func fact(x int) uint64 {
	for i := 1; i <= x; i++ {
		y *= uint64(i)
		fmt.Println("Il fattoriale di", i, "è", y)
	}
	return y
}

func main() {
	fmt.Print("Inserire numero non negativo: ")
	fmt.Scan(&x)
	
	for x < 0 {
		fmt.Print("Inserire numero NON negativo!: ")
		fmt.Scan(&x)
	}
	fmt.Println()
	fact(x)
	fmt.Println()
}
