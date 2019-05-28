package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	)
	
	var in string
	
func main() {
	//input string
	fmt.Print("Inserire frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in = scanner.Text()
		in += "\n"
	}
	
	lettere := make([]rune, 0, len(in))
	var maiusc bool
	for _, c := range in {
		if unicode.IsLetter(c) {
			maiusc = !maiusc
			if maiusc {
				c = unicode.ToUpper(c)
			} else if !maiusc {
				c = unicode.ToLower(c)
			}
		}
		lettere = append(lettere, c)
	}
	fmt.Print(string(lettere))
}
