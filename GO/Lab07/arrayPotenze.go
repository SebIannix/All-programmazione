package main

import (
	"fmt"
	"math"
)

var a1, a2, a3 [10]float64

func main() {
	var x float64
	
	for i := 0; i < len(a1); i++ {
		a1[i] = math.Exp2(x)
		x++
	}
	//fmt.Println(a1)
	
	a2 = a1
	x = 0
	for i := range a2 {
		fmt.Println(a2[i])
	}
	var c2, d2 [10]float64
	for i := 0; i < len(a2); i++ {
		c2[i] = math.Mod(a2[i], 10)
		d2[i] = c2[i] - a2[i]
		x++
		//fmt.Println(a2, "	", c2, "	", d2)
	}
}
