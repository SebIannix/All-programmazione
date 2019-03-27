package main

import (
	"fmt"
	)
	
	var in1, in2 int
	
//Check numeri primi
func isPrime(value int) bool {
	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
            }
	}
	return value > 1
}
	
func main() {
	//Input e catch numeri < 0
	fmt.Print("Inserire primo numero: ")
	fmt.Scan(&in1)
	for in1 <= 0 {
		fmt.Print(in1, " non è strettamente maggiore di 0. Inserire un altro numero: ")
		fmt.Scan(&in1)
		continue
	}
	
	fmt.Print("Inserire secondo numero: ")
	fmt.Scan(&in2)
	for in2 <= 0 {
		fmt.Print(in2, " non è strettamente maggiore di 0. Inserire un altro numero: ")
		fmt.Scan(&in2)
		continue
	}
	
	if isPrime(in1) {
		if isPrime(in2) {
			if in2 == (in1 + 2) {
				fmt.Println(in1, "e", in2, "sono primi gemelli.")
			} else {
				fmt.Println(in1, "e", in2, "non sono primi gemelli.")
			}
		}
	}
}
