package main

import (
	"fmt"
	"os"
	 "encoding/json"
)

func main() {
	nom := demanderNom()
	if nom == "" {
		donnees := map[string]any{
	"nom": nom,
	"mot": "",
	"lettres_jouées": []string{},
	"essaies_restants": 6,
		}
}

contenuJSON, err := json.MarshalIndent(donnees, "", "  ")
if err != nil {
	fmt.Println("Erreur de conversion JSON :", err)
	return
}

err = os.WriteFile("save.json", contenuJSON, 0644)
if err != nil {
	fmt.Println("Erreur de sauvegarde :", err)
	return
}
		return
	}
	fmt.Printf("Bonjour %s !\n", nom)

	if _, err := demarrerMot(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}
