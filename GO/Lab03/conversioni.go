package main

import "fmt"

func main() {
	var valore, out, Out float64
	var scelta int

	fmt.Print("Scegli la conversione:\n")
	fmt.Print("1: secondi -> ore\n")
	fmt.Print("2: secondi -> minuti\n")
	fmt.Print("3: minuti -> ore\n")
	fmt.Print("4: minuti -> secondi\n")
	fmt.Print("5: ore -> secondi\n")
	fmt.Print("6: ore -> minuti\n")
	fmt.Print("7: minuti -> giorni e ore\n")
	fmt.Print("8: minuti -> anni e giorni\n")

	fmt.Scan(&scelta)

	fmt.Print("Inserisci il valore da convertire: ")

	fmt.Scan(&valore)

	if scelta == 1 {
		out = valore / 3600
		fmt.Println("Ore:", out)
	} else if scelta == 2 {
		out = valore / 60
		fmt.Println("Minuti:", out)
	} else if scelta == 3 {
		out = valore / 60
		fmt.Println("Ore:", out)
	} else if scelta == 4 {
		out = valore * 60
		fmt.Println("Secondi:", out)
	} else if scelta == 5 {
		out = valore * 3600
		fmt.Println("Secondi:", out)
	} else if scelta == 6 {
		out = valore * 60
		fmt.Println("Minuti:", out)
	} else if scelta == 7 {
		out = valore / 60
		Out = valore / 1440
		fmt.Println("Giorni:", Out, ", e ore:", out)
	} else if scelta == 8 {
		Out = valore / 525600
		out = valore / 1440
		fmt.Println("Anni:", Out, ", e giorni:", out)
	}
}
