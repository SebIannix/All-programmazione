package main

import (
	"fmt"
	"strings"
	)

	var in int
	var ast string = "* "
	
func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&in)
	
	var line = strings.Repeat(ast, in)
	var linePiù = strings.Replace(line, "* ", "+ ", -1)

	for i := 1; i <= in; i++ {
		if i%2 == 0 {
			fmt.Println(linePiù)
		} else {
			fmt.Println(line)
		}
	}
}
