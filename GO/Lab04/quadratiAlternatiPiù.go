package main

import (
	"fmt"
	"strings"
	)

	var in int
	var più string = "+ "
	var ast string = "* "
	
func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&in)
	
	line		:= strings.Repeat(più, in-1)
	linePiù	:= strings.Repeat(più, in)
	linePiù2	:= strings.Replace(linePiù, "+ ", "* ", -1)
	lineMix	:= strings.Replace(line, "+ ", "* ", in-(in-1))
	lineMix2	:= strings.TrimRight(lineMix, " ")
	
	for i := 1; i <= in; i++ {
		if i != 1 && i != in {
			fmt.Println(lineMix2, "*")
		} else {
			fmt.Println(linePiù2)
		}
	}
	
}
