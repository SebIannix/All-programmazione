package main

import (
	"fmt"
	"time"
	)
	
	var nome string
	
	
func main() {
	fmt.Print("Come ti chiami?: ")
	fmt.Scan(&nome)
	
	data := time.Now()
	
	fmt.Print("Ciao ", nome, ", oggi è il ", data)
	fmt.Println()
}
