//unfinished
package main

import (
	"fmt"
)

var in, n, N int

func main() {
	fmt.Print("Inserire intero positivo: ")
	fmt.Scan(&in)	//input
	
	for in <= 0 {	//check in positivo
		fmt.Print("Inserire intero positivo valido: ")
		fmt.Scan(&in)
	}
	
	for i := 1; i <= in; i++ {	//max quadrato
		n = 1
		N = n*n
		for N <= in {
			n++
			N = n*n
		}
		
	}

	fmt.Println(N)
}
