package main

import (
	"fmt"
	)

	var num, div, numDiv, sumDiv int

func main() {
	fmt.Print("Inserire numero: ")
	fmt.Scan(&num)

	//Check divisori
	for div = 1; div < num; div++ {
		if num%div == 0 {
			numDiv++;
			sumDiv += div
		}
	}
		
	if sumDiv == num {
		fmt.Println(num, "è un numero perfetto.")
	} else {
		fmt.Println(num, "non è un numero perfetto.")
	}
		
}
