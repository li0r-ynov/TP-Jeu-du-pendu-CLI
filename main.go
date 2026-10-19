package main

import (
	"fmt"
	"os"
)

func main() {
	nom := demanderNom()
	if nom == "" {
		return
	}
	fmt.Printf("Bonjour %s !\n", nom)

	if _, err := demarrerMot(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}
