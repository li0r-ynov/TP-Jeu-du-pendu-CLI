package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "jouer" {
		fmt.Println("Utilisation : go run . jouer")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	nom := demanderNom(scanner)
	if nom == "" {
		return
	}

	partie := Partie{
		Nom:             nom,
		LettresJouees:   []string{},
		EssaiesRestants: 6,
	}
	if err := partie.selectionnerMot(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		return
	}

	fmt.Printf("Bonjour %s !\n", nom)
	if err := sauvegarderPartie(partie); err != nil {
		fmt.Println("Erreur de sauvegarde :", err)
		return
	}

	for partie.EssaiesRestants > 0 && !partie.motTrouve() {
		fmt.Printf("\nMot à deviner : %s\n", partie.afficherMot())
		fmt.Printf("Essais restants : %d\n", partie.EssaiesRestants)
		fmt.Print("Entrez une lettre : ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Erreur lors de la lecture de la lettre :", err)
			}
			return
		}

		saisie := strings.TrimSpace(scanner.Text())
		lettres := []rune(saisie)
		if len(lettres) != 1 || !unicode.IsLetter(lettres[0]) {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		lettre := unicode.ToLower(lettres[0])
		if partie.lettreDejaJouee(lettre) {
			fmt.Println("Vous avez déjà proposé cette lettre.")
			continue
		}

		partie.LettresJouees = append(partie.LettresJouees, string(lettre))

		lettrePresente := false
		for _, lettreMot := range partie.MotSecret {
			if unicode.ToLower(lettreMot) == lettre {
				lettrePresente = true
				break
			}
		}
		if !lettrePresente {
			partie.EssaiesRestants--
			fmt.Println("Cette lettre n'est pas dans le mot.")
		}

		if err := sauvegarderPartie(partie); err != nil {
			fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
			return
		}
	}

	if partie.motTrouve() {
		fmt.Printf("\nBravo %s, tu as trouvé le mot : %s\n", partie.Nom, partie.MotSecret)
	} else {
		fmt.Printf("\nPerdu ! Le mot était : %s\n", partie.MotSecret)
	}
}

func (partie Partie) afficherMot() string {
	var affichage strings.Builder
	for _, lettre := range partie.MotSecret {
		if !unicode.IsLetter(lettre) || partie.lettreDejaJouee(unicode.ToLower(lettre)) {
			affichage.WriteRune(lettre)
		} else {
			affichage.WriteRune('_')
		}
		affichage.WriteRune(' ')
	}
	return strings.TrimSpace(affichage.String())
}

func (partie Partie) motTrouve() bool {
	for _, lettre := range partie.MotSecret {
		if unicode.IsLetter(lettre) && !partie.lettreDejaJouee(unicode.ToLower(lettre)) {
			return false
		}
	}
	return true
}

func (partie Partie) lettreDejaJouee(lettre rune) bool {
	for _, lettreJouee := range partie.LettresJouees {
		if len([]rune(lettreJouee)) == 1 && unicode.ToLower([]rune(lettreJouee)[0]) == lettre {
			return true
		}
	}
	return false
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
