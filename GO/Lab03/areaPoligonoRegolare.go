package main

import (
	"fmt"
	"math"
)

func main() {
	var n, l, area float64
	fmt.Print("Enter, respectively, the number of sides and their length: ")
	fmt.Scan(&n, &l)
	area = (n * math.Pow(1, 2)) / (4 * math.Tan(math.Pi/n))
	fmt.Println(area)
}
