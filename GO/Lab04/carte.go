package main

import (
	"fmt"
	)
	
	var c rune = 127153
	
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(string(c))
		c++
	}
}
