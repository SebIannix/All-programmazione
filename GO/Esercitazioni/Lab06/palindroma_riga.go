package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
	)

func main() {
	var in string
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserire frase: ")
	if s.Scan() {
		in = s.Text()
	}
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		panic(err)
	}
	m := reg.ReplaceAllString(in, "")
	n := strings.ToLower(m)
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
