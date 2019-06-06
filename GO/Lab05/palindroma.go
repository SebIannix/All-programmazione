package main

import (
	"fmt"
)

var in string

func main() {
	fmt.Print("Inserire parola: ")
	fmt.Scan(&in)
	
	r := []rune(in)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	
	if in == string(r) {
		fmt.Println("Palindroma!")
	} else {
		fmt.Println("Non palindroma.")
	}
}
