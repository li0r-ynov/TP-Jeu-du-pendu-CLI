package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Partie contient tout ce qu'il faut pour reprendre une partie en cours.
type Partie struct {
	Pseudo         string   `json:"pseudo"`
	Mot            string   `json:"mot"`
	LettresJouees  []string `json:"lettres_jouees"`
	EssaisRestants int      `json:"essais_restants"`
}

const fichierSauvegarde = "save.json"
const fichierHistorique = "score.csv"

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
	if strings.TrimSpace(partie.Mot) == "" || strings.TrimSpace(partie.Pseudo) == "" {
		return Partie{}, errors.New("aucune partie complète à reprendre dans save.json")
	}
	if partie.LettresJouees == nil {
		partie.LettresJouees = []string{}
	}
	return partie, nil
}

// ajouterHistorique s'appelle à la fin d'une partie.
func ajouterHistorique(partie Partie, resultat string) error {
	fichier, err := os.OpenFile(fichierHistorique, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ouverture de %s : %w", fichierHistorique, err)
	}
	defer fichier.Close()

	info, err := fichier.Stat()
	if err != nil {
		return fmt.Errorf("lecture des informations de %s : %w", fichierHistorique, err)
	}
	writer := csv.NewWriter(fichier)
	if info.Size() == 0 {
		if err := writer.Write([]string{"date", "pseudo", "mot", "resultat", "essais_restants"}); err != nil {
			return fmt.Errorf("écriture de l'en-tête CSV : %w", err)
		}
	}
	if err := writer.Write([]string{
		time.Now().Format("2006-01-02 15:04:05"),
		partie.Pseudo,
		partie.Mot,
		resultat,
		fmt.Sprintf("%d", partie.EssaisRestants),
	}); err != nil {
		return fmt.Errorf("écriture du résultat CSV : %w", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("finalisation de l'historique : %w", err)
	}
	return nil
}
