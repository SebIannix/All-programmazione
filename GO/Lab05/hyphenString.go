package main

import (
	"fmt"
	"bytes"
	"bufio"
	"os"
)

var in string

func tratteggia(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n-1
	var l_1 = len(s)-1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('-')
		}
	}
	return buffer.String()
}

func main() {
	fmt.Print("Inserire frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		in = scanner.Text()
		in += "\n"
	}
	fmt.Println(tratteggia(in, 1))
}
