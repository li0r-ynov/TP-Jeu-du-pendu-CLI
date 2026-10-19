package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "jouer" {
		fmt.Println("Utilisation : go run . jouer")
		return
	}

	nom := demanderNom()
	if nom == "" {
		return
	}
	fmt.Printf("Bonjour %s !\n", nom)

	if _, err := demarrerMot(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}
