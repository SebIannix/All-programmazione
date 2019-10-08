package main

import (
	"fmt"
	"os"
	"bufio"
)

type Persona struct {
	nome, cognome, telefono string
	indirizzo Indirizzo
}
	type Indirizzo struct {
		via, civico, CAP, città string
	}

func AddContact(x Persona) Persona {
	var contatto Persona
	
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		in = s.Text()
		in += "\n"
	}
	
}

func main() {
	
}
