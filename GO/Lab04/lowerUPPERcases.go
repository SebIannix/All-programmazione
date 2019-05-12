package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	)

	var in string

func main() {
	fmt.Print("Introduci la riga di testo: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in = scanner.Text()
		in += "\n"
	}
	
	fmt.Print("Riga maiuscola: ", strings.ToUpper(in))
	fmt.Print("Riga minuscola: ", strings.ToLower(in))	
}
