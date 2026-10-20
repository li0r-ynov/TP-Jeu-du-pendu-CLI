package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

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
	// Migration des sauvegardes produites par l'ancienne version.
	if partie.Nom == "" {
		var ancienne struct {
			Pseudo  string   `json:"pseudo"`
			Mot     string   `json:"mot"`
			Lettres []string `json:"lettres_jouees"`
			Essais  int      `json:"essais_restants"`
		}
		if err := json.Unmarshal(contenu, &ancienne); err == nil {
			partie = Partie{Nom: ancienne.Pseudo, MotSecret: ancienne.Mot, LettresJouees: ancienne.Lettres, EssaiesRestants: ancienne.Essais}
		}
	}
	if strings.TrimSpace(partie.Nom) == "" || strings.TrimSpace(partie.MotSecret) == "" {
		return Partie{}, fmt.Errorf("aucune partie complète à reprendre dans save.json")
	}
	if partie.EssaiesRestants < 0 || partie.EssaiesRestants > 6 {
		return Partie{}, fmt.Errorf("nombre d'essais invalide dans save.json")
	}
	if partie.LettresJouees == nil {
		partie.LettresJouees = []string{}
	}
	return partie, nil
}

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
	w := csv.NewWriter(fichier)
	if info.Size() == 0 {
		if err := w.Write([]string{"date", "nom", "mot", "résultat", "essais_restants"}); err != nil {
			return fmt.Errorf("écriture de l'en-tête CSV : %w", err)
		}
	}
	if err := w.Write([]string{time.Now().Format("2006-01-02 15:04:05"), partie.Nom, partie.MotSecret, resultat, fmt.Sprintf("%d", partie.EssaiesRestants)}); err != nil {
		return fmt.Errorf("écriture de l'historique : %w", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("finalisation de l'historique : %w", err)
	}
	return nil
}
