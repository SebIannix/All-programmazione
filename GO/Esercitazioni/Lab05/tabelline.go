package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Print("Inserire numero positivo: ")
	fmt.Scan(&n)
	
	for n <= 0 {
		fmt.Print("Inserire numero positivo: ")
		fmt.Scan(&n)
	}
	
	//tabellina
	w := 1
	for i := 1; i <= n; i++ {
		o := n*i
		fmt.Printf("%v * %v = %v\n", n, w, o)
		w++
	}
}
