package main

import (
	"fmt"
	"math/rand"
	"time"
)

var a[5][5]int

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := (i*j+1)*100
			a[i] = rand.Intn(x)
			a[j] = rand.Intn(x)
			fmt.Printf("Array[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
