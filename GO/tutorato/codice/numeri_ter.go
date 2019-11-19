package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	
	somma := 0
	d := 0
	
		for n != 0 {
			somma += n
			d++
			fmt.Scan(&n)
		}
	if d == 0 {
		fmt.Println("Nessun numero valido inserito.")
	} else {
	fmt.Println("Media:", somma/d)
	}
}
