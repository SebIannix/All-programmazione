package main

import (
	"fmt"
)
	
var in, sum, numLetti int

func main() {
	fmt.Println("	--> Inserire numeri tra 1 e 9; inserire 0 per terminare il programma:")
	for {
		fmt.Scan(&in)
		if in > 0 && in < 10 {
			sum += in
			numLetti++
			numStored := make([]int, numLetti)
			for v := range numStored {
				numStored[v] = in
				
				}
			fmt.Println("Slice:", numStored)
			}
		if in == 0 {
			break
		}

	}
	
	fmt.Println("Numeri letti:", numLetti)
	fmt.Println("Somma finale dei numeri leciti:", sum)

}
