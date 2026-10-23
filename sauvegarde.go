package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
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
	if partie.LettresJouees == nil {
		var ancienChamp struct {
			Lettres []string `json:"lettres_jouées"`
		}
		if err := json.Unmarshal(contenu, &ancienChamp); err == nil {
			partie.LettresJouees = ancienChamp.Lettres
		}
	}
	if strings.TrimSpace(partie.Nom) == "" || !WordIsValid(partie.MotSecret) {
		return Partie{}, fmt.Errorf("aucune partie complète à reprendre dans save.json")
	}
	if partie.EssaiesRestants < 0 || partie.EssaiesRestants > 6 {
		return Partie{}, fmt.Errorf("nombre d'essais invalide dans save.json")
	}
	if partie.LettresJouees == nil {
		partie.LettresJouees = []string{}
	}
	for _, lettre := range partie.LettresJouees {
		runes := []rune(lettre)
		if len(runes) != 1 || !unicode.IsLetter(runes[0]) {
			return Partie{}, fmt.Errorf("lettre invalide dans save.json")
		}
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

func afficherHistorique() error {
	fichier, err := os.Open(fichierHistorique)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Aucune partie terminée dans l'historique.")
			return nil
		}
		return fmt.Errorf("lecture de %s : %w", fichierHistorique, err)
	}
	defer fichier.Close()

	lecteur := csv.NewReader(fichier)
	lecteur.FieldsPerRecord = -1
	for ligne := 0; ; ligne++ {
		record, err := lecteur.Read()
		if err != nil {
			if err == io.EOF {
				if ligne == 0 {
					fmt.Println("Aucune partie terminée dans l'historique.")
				}
				return nil
			}
			return fmt.Errorf("lecture de l'historique : %w", err)
		}
		if ligne == 0 {
			continue // première ligne : noms des colonnes
		}
		if len(record) < 5 {
			return fmt.Errorf("ligne %d de score.csv incomplète", ligne+1)
		}
		fmt.Printf("%s | %s | mot : %s | résultat : %s | essais restants : %s\n", record[0], record[1], record[2], record[3], record[4])
	}
}
