package main

import (
	"fmt"
	)
	
	var numero int
	
func main() {
	fmt.Print("Inserire numero: ")
	fmt.Scan(&numero)
	
	fmt.Println("1 x", numero, "=", 1*numero)
	fmt.Println("2 x", numero, "=", 2*numero)
	fmt.Println("3 x", numero, "=", 3*numero)
	fmt.Println("4 x", numero, "=", 4*numero)
	fmt.Println("5 x", numero, "=", 5*numero)
	fmt.Println("6 x", numero, "=", 6*numero)
	fmt.Println("7 x", numero, "=", 7*numero)
	fmt.Println("8 x", numero, "=", 8*numero)
	fmt.Println("9 x", numero, "=", 9*numero)
	fmt.Println("10 x", numero, "=", 10*numero)
}
