package main

import (
	"fmt"
	"math"
	)

var x, y, max, min float64

func main() {
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	fmt.Println("Sum:", x+y)
	fmt.Println("Difference:", max-min)
	fmt.Println("Division:", x/y)
	fmt.Println("Multiplication:", x*y)
	fmt.Println("Power:", math.Pow(x, y))
	fmt.Println("Average:", (x+y)/2)
}
