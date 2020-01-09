package main

import (
	"fmt"
	"os"
	"bufio"
	)

func main() {
	var in string
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		in = s.Text()
	}
	switch in {
		case "Z":
			fmt.Println("*****")
			fmt.Println("   *")
			fmt.Println("  *")
			fmt.Println(" *")
			fmt.Println("*****")
		case "T":
			fmt.Println("*****")
			fmt.Println("  *  ")
			fmt.Println("  *  ")
			fmt.Println("  *  ")
			fmt.Println("  *  ")
		case "L":
			fmt.Println("*")
			fmt.Println("*")
			fmt.Println("*")
			fmt.Println("*")
			fmt.Println("*****")
		default:
			fmt.Println("Input non valido.")
	}
}
