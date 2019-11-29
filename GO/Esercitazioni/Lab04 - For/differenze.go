package main

import (
	"fmt"
)

func main() {
	var in []int
	fmt.Println("Inserire valori:")
	
	for i := 0; ; i++ {
		in = append(in, 0)
		_, err := fmt.Scanf("%d", &in[i])	//input ints
		if err != nil {
			in = in[:len(in)-1]
			break
		}
	}
	
	var a int
	diff := 0
	for _, i := range in {	//rassegna valori
		diff = i - a	//differenza
		fmt.Printf("Differenza tra %v e %v: %v\n", i, a, diff)
		a = i		//assegnazione valore precedente
	}
}
