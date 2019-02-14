package main

import "fmt"

func main() {
	var x, y, min, max int
	fmt.Print("Enter 2 ints separated by whitespace: \n")
	fmt.Scan(&x, &y)
	if x > y {
		fmt.Println("Max = ", x, "\n----------\n")
	} else {
		fmt.Print("Max = ", y, "\n----------")
	}
	fmt.Print("Enter 2 more ints: ")
	fmt.Scan(&min, &max)
	if min > max {
		min, max = max, min
	}
	fmt.Println(min, " <= ", max)
}
