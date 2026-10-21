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

fmt.Printf("Bienvenue %s !\n", nom)

_, err := demarrerMot()
if err != nil {
    fmt.Println("Erreur :", err)
    return
}
}
