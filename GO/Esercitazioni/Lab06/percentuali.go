package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	)

func PercentOf(voti int, tot int) float64 {
	percent := (float64(voti) * float64(100)) / float64(tot)
	return percent
}

func main() {
	var suf, ins []int
	var in string
	var tot, totSuf, totIns int = 0, 0, 0
	fmt.Print("Inserire sequenza di interi: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = scanner.Text()
		//Stop storing ints when typing -1
		i, err := strconv.Atoi(in)
		if err == nil && i == -1 {
			break
		}
		//Store ints in []num between 0-30
		for _, valori := range strings.Fields(in) {
			n, err := strconv.Atoi(valori)
			if err == nil && n >= 0 && n <= 30 {
				if n > 17 {
					suf = append(suf, n)
					tot++
					totSuf++
				}
				if n <= 17 {
					ins = append(ins, n)
					tot++
					totIns++
				}
			}
		}
	}
	fmt.Println("Sufficienze:", PercentOf(totSuf, tot), "%")
	fmt.Println("Insufficienze:", PercentOf(totIns, tot), "%")
}
