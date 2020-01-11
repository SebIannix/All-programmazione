/*
Frazioni
========

Scrivere un programma in Go (il file deve chiamarsi 'frazioni.go') dotato di:


- una funzione
	minTerm(n, d int) (int, int, error)
che, dati due interi, n1 e d1 che rappresentano il numeratore
e denominatore di una frazione, restituisce numeratore e
denominatore della frazione stessa ridotta ai minimi termini (*)
equivalente a n/d e un errore "divisore nullo" nel caso d sia
uguale a 0.
Nota: per creare una variabile di tipo errore e inizializzarla, usare
la funzione New del package errors, passando come argomento la stringa opportuna.
Ad esempio:
errors.New("divisore nullo")


- una funzione
	sum(n1, d1, n2, d2 int) (int, int)
che restituisce numeratore e denominatore della frazione somma di n1/d1 e n2/d2 ridotta ai minimi termini

- una funzione
	main()
che legge 4 valori interi non negativi da linea di comando (n1, d1, n2, d2), che rappresentano le due frazioni n1/d1 e n2/d2, ed emette su standard output:

	- il messaggio di errore "numero di parametri non valido" se il numero di parametri interi è diverso da 4

	- un messaggio di errore se una o tutte e due le frazioni hanno denominatore nullo

	- altrimenti la loro somma ridotta ai minimi termini

Si può assumere che i parametri passati da linea di comando siano tutti numeri di tipo int e non negativi.

(*) Suggerimento: Un modo (non efficiente ma semplice) di
ridurre una frazione n/d ai minimi termini è dividere
via via sia n sia d per ciascun divisore di entrambi n e d.



Esempi
======

1) eseguendo:

$ go run frazioni.go 1 2 1 4

si ottiene:

3 / 4


2) eseguendo:

$ go run frazioni.go 4 12 2 0

si ottiene:

divisore nullo


3) eseguendo:

$ go run frazioni.go 3 10 19 5 7

si ottiene:

numero di parametri non valido


4) eseguendo:

$ go run frazioni.go 456 342 78 98

si ottiene:

313 / 147


5) eseguendo:

$ go run frazioni.go 4 12 2 3

si ottiene:

1 / 1

*/
package main

import (
	"fmt"
	"os/exec"

	//"strings"
	//"log"
	"testing"
)

func TestMainMultiplo(t *testing.T) {
	lanciaEcontrolla("1", "20", "1", "30", "1 / 12\n", t)
	lanciaEcontrolla("1", "2000", "1", "2000", "1 / 1000\n", t)
	lanciaEcontrolla("1", "2", "1", "4", "3 / 4\n", t)
	lanciaEcontrolla("4", "12", "2", "0", "divisore nullo\n", t)
	lanciaEcontrolla("456", "342", "78", "98", "313 / 147\n", t)
	lanciaEcontrolla("4", "12", "2", "3", "1 / 1\n", t)
}

func lanciaEcontrolla(n1, d1, n2, d2, out string, t *testing.T) {
	fmt.Println("----main-------------------------------")
	subproc := exec.Command("./frazioni", n1, d1, n2, d2) // presuppone che sia già stato compilato
	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}
	fmt.Printf("Output:\n%s\n", string(stdout))
	fmt.Printf("Expected output:\n%s\n", out)
	if string(stdout) != out {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
	subproc.Process.Kill()
}

func TestNumParSbagliato(t *testing.T) {
	fmt.Println("-----main par sbagliati------------------------------")
	subproc := exec.Command("./frazioni", "1", "2", "3", "4", "5") // presuppone che sia già stato compilato
	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}
	fmt.Printf("Output:\n%s\n", string(stdout))

	out := "numero di parametri non valido\n"
	fmt.Printf("Expected output:\n%s\n", out)
	if string(stdout) != out {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
	subproc.Process.Kill()
}

// giusto per testare se la funzione c'è
func TestMinTerm(t *testing.T) {
	fmt.Println("-----min term------------------------------")

	n, d, err := minTerm(5, 2)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}

	if n != 5 || d != 2 {
		fmt.Println(">>> minTerm FAIL!")
		t.Fail()
	}
	fmt.Println(n, d)
}

func TestSum(t *testing.T) {
	fmt.Println("-----sum------------------------------")

	ns, ds := sum(2, 5, 3, 5)

	if (ns != 1) || (ds != 1) {
		fmt.Println(">>> sum FAIL!")
		t.Fail()
	}
}
