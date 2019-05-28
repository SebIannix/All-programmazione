package main

import (
	"fmt"
	"bufio"
	"os"
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
	
	//counter vocali maiusc
	VOCALI := 0
	for _, value := range in {
		switch value {
			case 'A', 'E', 'I', 'O', 'U':
				VOCALI++
		}
	}
	
	//counter consonanti maiusc
	CONSONANTI := 0
	for _, value := range in {
		switch value {
			case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
				CONSONANTI++
		}
	}
	
	//counter vocali minusc
	vocali := 0
	for _, value := range in {
		switch value {
			case 'a', 'e', 'i', 'o', 'u':
				vocali++
		}
	}
	
	//counter consonanti minusc
	consonanti := 0
	for _, value := range in {
		switch value {
			case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
				consonanti++
		}
	}
	
	//outputs
	fmt.Printf("N° vocali maiuscole:		%d", VOCALI)
	fmt.Println()
	fmt.Printf("N° vocali minuscole:		%d", vocali)
	fmt.Println()
	fmt.Printf("N° consonanti maiuscole:	%d", CONSONANTI)
	fmt.Println()
	fmt.Printf("N° consonanti minuscole:	%d", consonanti)
	fmt.Println()
}
