package main

import (
	"fmt"
	)
	
	var in int
	
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
	fmt.Print("Inserire numero: ")
	fmt.Scan(&in)
	
	//Catch dei numeri <= 0
	for in <= 0 {
		fmt.Print(in, " non è strettamente maggiore di 0. Inserire un altro numero: ")
		fmt.Scan(&in)
		continue
	}
	
	fmt.Println(isPrime(in))
	
}
