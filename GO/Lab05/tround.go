package main

import (
	"fmt"
	"math"
	)

	var metodo int
	var in float64
	var out float64
	var depth int
	
func main() {
	//Selezione metodo
	fmt.Println("Seleziona metodo:")
	fmt.Println("	1 - Troncamento")
	fmt.Println("	2 - Arrotondamento")
	fmt.Scan(&metodo)
	
	//Catch: in esclusi
	for metodo != 1 && metodo != 2 {
		fmt.Print("Scegliere tra le due opzioni: ")
		fmt.Scan(&metodo)
	}
	
	//input numero
	fmt.Print("Inserire numero reale: ")
	fmt.Scan(&in)
	
	//input precisione
	fmt.Print("Definire quantità di cifre dopo la virgola: ")
	fmt.Scan(&depth)
	
	//Catch: precisione impossibile
	for depth <= 0 {
		fmt.Print("Definire correttamente la quantità di cifre DOPO la virgola: ")
		fmt.Scan(&depth)
	}
	
	//calcolo precisione
	var depthConv = math.Pow10(depth)
	
	//rounding
	if metodo == 2 {
		inMult	:= in*depthConv
		inRound	:= math.Round(inMult)
		out		= inRound/depthConv
	//troncamento
	} else if metodo == 1 {
		inMult	:= in*depthConv
		inInt	:= int64(inMult)
		inTrunc := float64(inInt)
		out		= inTrunc/depthConv
	}
	
	//output
	fmt.Println(out)
}
