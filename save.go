package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func gererCommandeSauvegarde(partie Partie, saisie string) (bool, error) {
	if saisie != "!save" {
		return false, nil
	}

	if err := sauvegarderPartie(partie); err != nil {
		return false, err
	}
	return true, nil
}

func sauvegarderPartie(partie Partie) error {
	contenuJSON, err := json.MarshalIndent(partie, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible de convertir la partie en JSON : %w", err)
	}
	if err := os.WriteFile("save.json", contenuJSON, 0644); err != nil {
		return fmt.Errorf("impossible d'enregistrer la partie : %w", err)
	}
	return nil
}

func chargerPartie() (Partie, error) {
	contenuJSON, err := os.ReadFile("save.json")
	if err != nil {
		return Partie{}, fmt.Errorf("impossible de lire save.json : %w", err)
	}

	var partie Partie
	if err := json.Unmarshal(contenuJSON, &partie); err != nil {
		return Partie{}, fmt.Errorf("impossible de lire la sauvegarde : %w", err)
	}
	if partie.Nom == "" || partie.MotSecret == "" {
		return Partie{}, fmt.Errorf("la sauvegarde ne contient pas de partie valide")
	}

	return partie, nil
}
