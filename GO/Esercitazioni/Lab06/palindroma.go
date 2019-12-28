package main

import (
	"fmt"
	"os"
	"strings"
	)

func main() {
	in := os.Args[1]
	fmt.Println("Parola:", in)
	n := strings.ToLower(in)
	r := []rune(n)
	
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	
	if n == string(r) {
		fmt.Println("Palindroma.")
	} else {
		fmt.Println("Non palindroma.")
	}
}
