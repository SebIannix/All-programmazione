package main

import "fmt"

var in int
var c, f float64

func main() {
	fmt.Println("Select conversion: C°-> F° (1) or F°-> C° (2)")
	fmt.Scan(&in)
	if in == 1 {
		fmt.Print("Enter C° temperature: ")
		fmt.Scan(&c)
		f = c * 9.0/5.0 + 32
		fmt.Println("F° temperature:", f)
	} else if in == 2 {
		fmt.Print("Enter F° temperature: ")
		fmt.Scan(&f)
		c = (f - 32) * 5.0/9.0
		fmt.Println("C° temperature:", c)
	}
}
