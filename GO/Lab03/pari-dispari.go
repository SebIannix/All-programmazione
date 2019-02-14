package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter an int: ")
	fmt.Scan(&x)
	if x % 2 == 0 {
		fmt.Println("Number is even.")
	} else {
		fmt.Println("Number is odd.")
	}
}
