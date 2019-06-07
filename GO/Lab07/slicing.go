package main

import (
	"fmt"
	"math"
)

var a1 [10]float64

func main() {
	//Array
	var x float64
	for i := 0; i < len(a1); i++ {
		a1[i] = math.Exp2(x)
		x++
	}
	fmt.Println("Array:\t", a1)
	
	//Slices init
	var s1, s3 []float64
	
	//Slice 1
	s1 = a1[4:7]
	fmt.Println("Original slice:\t", s1)
	s1 = a1[4:10]
	fmt.Println("Expanded slice:\t", s1)
	
	//Slice 2
	s3 = a1[0:10]
	for i, j := 0, len(s3)-1; i < j; i, j = i+1, j-1 {
		s3[i], s3[j] = s3[j], s3[i]
	}
	fmt.Println("Reverse slice:\t", s3)
}
