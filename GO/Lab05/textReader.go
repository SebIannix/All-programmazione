package main

import(
	"fmt"
	"bufio"
	"os"
	"unicode"
)

var in string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var righe []string
	for scanner.Scan() {
		in = scanner.Text()
		if len(in) == 1 {
			if in[0] == '\x1D' {
				break
			}
		}
		righe = append(righe, in)
	}
	if len(righe) > 0 {
	fmt.Println()
		for _, in := range righe {
			fmt.Println(in)
		}
	}
	
	var count int
	for _, r := range in {
		if unicode.IsLetter(r) {
			count++
		}
	}
	fmt.Println(count)
}
