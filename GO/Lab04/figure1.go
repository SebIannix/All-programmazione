package main

import (
	"fmt"
	//"strings"
	)

	var in int
	var ast string = "*"
	var più string = "+"
	
func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&in)

	for i := 1; i <= in; i++ {
		//line := ast + più
		//lineOut := strings.Repeat(line, in)
		fmt.Print(ast)
		fmt.Print(più)
		
	}
	fmt.Println()
	
	/*
	var line = strings.Repeat(ast, in)
	var linePiù = strings.Replace(line, "* ", "+ ", -1)

	for i := 1; i <= in; i++ {
		if i%2 == 0 {
			fmt.Println(linePiù)
		} else {
			fmt.Println(line)
		}
	}
	*/
}
