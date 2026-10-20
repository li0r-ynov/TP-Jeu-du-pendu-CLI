package main

import (
	"encoding/json"
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

	donnees := map[string]any{
		"nom":              nom,
		"mot":              "",
		"lettres_jouées":   []string{},
		"essaies_restants": 6,
	}

	contenuJSON, err := json.MarshalIndent(donnees, "", "  ")
	if err != nil {
		fmt.Println("Erreur de conversion JSON :", err)
		return
	}

	if err := os.WriteFile("save.json", contenuJSON, 0644); err != nil {
		fmt.Println("Erreur de sauvegarde :", err)
		return
	}

	fmt.Printf("Bonjour %s !\n", nom)

	if _, err := demarrerMot(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}
