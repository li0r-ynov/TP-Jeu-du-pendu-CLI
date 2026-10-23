package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

const fichierHistorique = "score.csv"

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
		if err := writer.Write([]string{"date", "nom", "mot", "résultat", "essais_restants"}); err != nil {
			return fmt.Errorf("écriture de l'en-tête CSV : %w", err)
		}
	}
	if err := writer.Write([]string{
		time.Now().Format("2006-01-02 15:04:05"),
		partie.Nom,
		partie.MotSecret,
		resultat,
		fmt.Sprintf("%d", partie.EssaiesRestants),
	}); err != nil {
		return fmt.Errorf("écriture de l'historique : %w", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
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

	reader := csv.NewReader(fichier)
	reader.FieldsPerRecord = -1
	for ligne := 0; ; ligne++ {
		record, err := reader.Read()
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
			continue
		}
		if len(record) < 5 {
			return fmt.Errorf("ligne %d de score.csv incomplète", ligne+1)
		}
		fmt.Printf("%s | %s | mot : %s | résultat : %s | essais restants : %s\n",
			record[0], record[1], record[2], record[3], record[4])
	}
}
