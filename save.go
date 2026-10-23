package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const fichierSauvegarde = "save.json"

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
	contenu, err := json.MarshalIndent(partie, "", "  ")
	if err != nil {
		return fmt.Errorf("conversion de la sauvegarde : %w", err)
	}
	if err := os.WriteFile(fichierSauvegarde, contenu, 0644); err != nil {
		return fmt.Errorf("écriture de %s : %w", fichierSauvegarde, err)
	}
	return nil
}

func chargerPartie() (Partie, error) {
	contenu, err := os.ReadFile(fichierSauvegarde)
	if err != nil {
		return Partie{}, fmt.Errorf("lecture de %s : %w", fichierSauvegarde, err)
	}

	var partie Partie
	if err := json.Unmarshal(contenu, &partie); err != nil {
		return Partie{}, fmt.Errorf("sauvegarde JSON invalide : %w", err)
	}

	var ancienne struct {
		Pseudo              string   `json:"pseudo"`
		Nom                 string   `json:"nom"`
		Mot                 string   `json:"mot"`
		LettresJouees       []string `json:"lettres_jouees"`
		LettresJoueesAccent []string `json:"lettres_jouées"`
		EssaisRestants      *int     `json:"essais_restants"`
	}
	if err := json.Unmarshal(contenu, &ancienne); err != nil {
		return Partie{}, fmt.Errorf("sauvegarde JSON invalide : %w", err)
	}
	if partie.Nom == "" {
		partie.Nom = ancienne.Nom
		if partie.Nom == "" {
			partie.Nom = ancienne.Pseudo
		}
	}
	if partie.MotSecret == "" {
		partie.MotSecret = ancienne.Mot
	}
	if partie.LettresJouees == nil {
		partie.LettresJouees = ancienne.LettresJouees
		if partie.LettresJouees == nil {
			partie.LettresJouees = ancienne.LettresJoueesAccent
		}
	}
	if ancienne.EssaisRestants != nil {
		partie.EssaiesRestants = *ancienne.EssaisRestants
	}

	if strings.TrimSpace(partie.Nom) == "" || !WordIsValid(partie.MotSecret) {
		return Partie{}, fmt.Errorf("aucune partie complète à reprendre dans %s", fichierSauvegarde)
	}
	if partie.EssaiesRestants < 0 || partie.EssaiesRestants > 6 {
		return Partie{}, fmt.Errorf("nombre d'essais invalide dans %s", fichierSauvegarde)
	}
	if partie.LettresJouees == nil {
		partie.LettresJouees = []string{}
	}
	for _, lettre := range partie.LettresJouees {
		runes := []rune(lettre)
		if len(runes) != 1 || !unicode.IsLetter(runes[0]) {
			return Partie{}, fmt.Errorf("lettre invalide dans %s", fichierSauvegarde)
		}
	}
	if partie.Difficulte == "" {
		partie.Difficulte = difficulteMot(partie.MotSecret)
	}
	if partie.Difficulte != "facile" && partie.Difficulte != "difficile" {
		return Partie{}, fmt.Errorf("difficulté invalide dans %s", fichierSauvegarde)
	}
	return partie, nil
}
