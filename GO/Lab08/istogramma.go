package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
)

func freq(s string) string {
	var r rune
	var v string
	m := make(map[rune]int)
	for _, r = range s {
		if unicode.IsSpace(r) == false {
			m[r]++
		}
	}
	for k, i := range m {
		fmt.Print(string(k), ": ", i, "\n")
	}
	return v
}

func main() {
	var in, out string
	
	fmt.Print("Inserire testo: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = scanner.Text()
		if in == "" {
			break
		}
		out += in
		out = strings.ToLower(out)
	}
	fmt.Println(freq(out))
}
