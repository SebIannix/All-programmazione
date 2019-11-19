package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Println("Inserisci tre numeri separati da spazi:")
	fmt.Scanln(&a, &b, &c)
	
	if a < b+c && b < a+c && c < a+b {
		if a == b && a == c {
			fmt.Println("Il triangolo è equilatero.")
		} else if a == b || a == c || b == c {
			fmt.Println("Il triangolo è isoscele.")
		} else {
			fmt.Println("Il triangolo è scaleno.")
		}
	} else {
		fmt.Println("Un lato non è minore della somma degli altri due.")
		fmt.Println("A:", a, "	B+C:", b+c)
		fmt.Println("B:", b, "	A+C:", a+c)
		fmt.Println("C:", c, "	A+B:", a+b)
	}
}
