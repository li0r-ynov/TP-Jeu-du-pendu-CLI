package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Utilisation : go run . jouer ou go run . reprendre")
		return
	}

	switch os.Args[1] {
	case "jouer":
		nouvellePartie()
	case "reprendre":
		reprendrePartie()
	default:
		fmt.Println("Commande inconnue. Utilise : go run . jouer ou go run . reprendre")
	}
}

func nouvellePartie() {
	nom := demanderNom()
	if nom == "" {
		return
	}

	mot, err := demarrerMot()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		return
	}

	partie := Partie{
		Pseudo:         nom,
		Mot:            mot,
		LettresJouees:  []string{},
		EssaisRestants: 6,
	}
	if err := sauvegarderPartie(partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
		return
	}
	fmt.Printf("Bonjour %s ! La partie est sauvegardée.\n", nom)
}

func reprendrePartie() {
	partie, err := chargerPartie()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Impossible de reprendre :", err)
		return
	}
	fmt.Printf("Sauvegarde chargée pour %s (%d essais restants).\n", partie.Pseudo, partie.EssaisRestants)
	fmt.Println("La logique de saisie des lettres doit maintenant reprendre avec cette partie.")
}
