package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Lanci(risultati, intervallo int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < risultati; i++ {
		fmt.Print(rand.Intn(intervallo), " ")
	}
	return risultati
}

func Gioco(x int) int {
	giocatori := "12345"
	for i := 0; i < len(giocatori); i++ {
		fmt.Print("Giocatore: ")
		Lanci(2, 7)
		fmt.Println()
		for j := 0; j < turni; j++ {
			
		}
	}
	return x
}

var turni, x int

func main() {
	fmt.Print("	Inserire numero di turni di gioco: ")
	fmt.Scan(&turni)
	fmt.Println(" --------- Lancio dadi in corso... ---------")
	
	Gioco(x)
}
