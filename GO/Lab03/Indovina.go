/*L'utente deve inserire un numero da tastiera per indovinare il numero fissato nel codice*/
package main

import "fmt"

func main() {
	var x, ind int = 0, 10
	fmt.Println("Prova a indovinare il numero fissato in codice.")
	fmt.Scan(&x)
	if x == ind {
		fmt.Println("Indovinato!")
	} else {
		if x < ind {
			fmt.Println("Troppo piccolo.")
		} else {
			fmt.Println("Troppo grande.")
		}
	}
}
