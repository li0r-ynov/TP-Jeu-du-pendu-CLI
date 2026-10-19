package main

import "fmt"

func main() {
	nom := demanderNom()
	if nom == "" {
		return
	}
	fmt.Printf("Bonjour %s !\n", nom)
}
