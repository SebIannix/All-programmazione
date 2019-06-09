package main

import (
	"fmt"
	"math/rand"
	"time"
)

var soglia int

func main() {
	fmt.Print("Inserire soglia limite: ")
	fmt.Scan(&soglia)
	
	rand.Seed(time.Now().UnixNano())
	
	var s []int
	
	for i := 0; i < 100; i++ {
		x := rand.Intn(100)
		if x <= soglia {
			break
		}
		s = append(s, x)
	}
	fmt.Print("Array: ", s, " ")
	fmt.Println()
}
