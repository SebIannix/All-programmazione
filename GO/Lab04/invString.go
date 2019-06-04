package main

import (
	"fmt"
	"bufio"
	"os"
)

var uIn string

func Reverse(uIn string) string {
	runes := []rune(uIn)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Introduci testo:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		if in == "" {
			fmt.Println("Lettura terminata.")
			break
		} else {
			uIn += in + "\n"
		}
	}
	
	fmt.Println("Testo invertito:")
	fmt.Println(Reverse(uIn))
}
