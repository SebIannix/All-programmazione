package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PercentOf(numeri int, tot int) int {
	percent := (numeri * 100) / tot
	return percent
}

func main() {
	var in int
	var uno, due, tre, quattro, cinque, sei, tot int = 0, 0, 0, 0, 0, 0, 0
	fmt.Print("Quanti lanci di dado fare?: ")
	fmt.Scan(&in)
	//Dice roll
	var dice = []int{1, 2, 3, 4, 5, 6}
	rand.Seed(time.Now().Unix())
	for i := 0; i < in; i++ {
	roll := rand.Intn(len(dice))
		switch roll {
			case 0:
				uno++
				tot++
			case 1:
				due++
				tot++
			case 2:
				tre++
				tot++
			case 3:
				quattro++
				tot++
			case 4:
				cinque++
				tot++
			case 5:
				sei++
				tot++
			}
		}
		fmt.Println("1:", uno, "-", PercentOf(uno, tot), "%")
		fmt.Println("2:", due, "-", PercentOf(due, tot), "%")
		fmt.Println("3:", tre, "-", PercentOf(tre, tot), "%")
		fmt.Println("4:", quattro, "-", PercentOf(quattro, tot), "%")
		fmt.Println("5:", cinque, "-", PercentOf(cinque, tot), "%")
		fmt.Println("6:", sei, "-", PercentOf(sei, tot), "%")
}
