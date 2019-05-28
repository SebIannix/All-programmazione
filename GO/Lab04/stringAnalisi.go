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
	
	//counter vowels
	vowels := 0
	for _, value := range in {
		switch value {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				vowels++
		}
	}
	
	//counter consonanti
	consonanti := 0
	for _, value := range in {
		switch value {
			case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
				consonanti++
		}
	}
	
	//counter maiusc
	maiusc := 0
	for _, value := range in {
		switch value {
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				maiusc++
		}
	}
	
	//counter minusc
	minusc := 0
	for _, value := range in {
		switch value {
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				minusc++
		}
	}
	
	//outputs
	fmt.Printf("N° lettere maiuscole:	%d", maiusc)
	fmt.Println()
	fmt.Printf("N° lettere minuscole:	%d", minusc)
	fmt.Println()
	fmt.Printf("N° vocali:		%d", vowels)
	fmt.Println()
	fmt.Printf("N° consonanti:		%d", consonanti)
	fmt.Println()
}
