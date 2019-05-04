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

	for i := 1; i <= in; i++ {
		fmt.Println(line)
	}
}
