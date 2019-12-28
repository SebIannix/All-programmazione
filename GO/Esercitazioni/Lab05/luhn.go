package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	)

func main() {
	//Check for args
	var in string
	if len(os.Args) <= 1 {
		fmt.Println("Sequenza vuota.")
		os.Exit(0)
	} else {
		in = os.Args[1]
	}
	var o []string = strings.Split(in, "")
//-----------------------------------------------
	//copying []string elements to []int
	var a []int
	for _, v := range o {
		w, err := strconv.Atoi(v)
		if err != nil {	//voids "multiple-value in single-value context" error
			panic(err)
		}
		a = append(a, w)
	}
//-----------------------------------------------
	//Ops on []int
	sum := 0
	for i := range a {
		if i % 2 == 0 {
			a[i] *= 2
			if a[i] > 9 {
				a[i] -= 9
			}
		}
		sum += a[i]
	}
	fmt.Println("Algoritmo Luhn:", a)
	fmt.Println("Somma:", sum)
	if sum % 10 == 0 {
		fmt.Println("Numero valido.")
	} else {
		fmt.Println("Numero non valido.")
	}
}
