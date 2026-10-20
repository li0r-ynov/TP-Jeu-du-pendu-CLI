package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 || (os.Args[1] != "jouer" && os.Args[1] != "reprendre") {
		fmt.Println("Utilisation : go run . jouer | reprendre")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	var partie Partie
	if os.Args[1] == "reprendre" {
		var err error
		partie, err = chargerPartie()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Impossible de reprendre :", err)
			return
		}
		fmt.Printf("Reprise de la partie de %s.\n", partie.Nom)
	} else {
		nom := demanderNom(scanner)
		if nom == "" {
			return
		}
		partie = Partie{Nom: nom, LettresJouees: []string{}, EssaiesRestants: 6}
		if err := partie.selectionnerMot(); err != nil {
			fmt.Fprintln(os.Stderr, "Erreur :", err)
			return
		}
	}
	if partie.EssaiesRestants <= 0 || partie.motTrouve() {
		fmt.Println("La partie sauvegardée est déjà terminée.")
		return
	}
	if err := sauvegarderPartie(partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
		return
	}
	if err := jouer(scanner, &partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}

func jouer(scanner *bufio.Scanner, partie *Partie) error {
	for partie.EssaiesRestants > 0 && !partie.motTrouve() {
		fmt.Printf("\nMot à deviner : %s\nEssais restants : %d\nEntrez une lettre : ", partie.afficherMot(), partie.EssaiesRestants)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("lecture de la lettre : %w", err)
			}
			fmt.Println("\nPartie interrompue et sauvegardée. Relancez avec `go run . reprendre`.")
			return sauvegarderPartie(*partie)
		}
		lettres := []rune(strings.TrimSpace(scanner.Text()))
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
		present := false
		for _, r := range partie.MotSecret {
			if unicode.ToLower(r) == lettre {
				present = true
				break
			}
		}
		if !present {
			partie.EssaiesRestants--
			fmt.Println("Cette lettre n'est pas dans le mot.")
		}
		if err := sauvegarderPartie(*partie); err != nil {
			return err
		}
	}
	resultat := "perdu"
	if partie.motTrouve() {
		resultat = "gagné"
		fmt.Printf("\nBravo %s, tu as trouvé le mot : %s\n", partie.Nom, partie.MotSecret)
	} else {
		fmt.Printf("\nPerdu ! Le mot était : %s\n", partie.MotSecret)
	}
	if err := sauvegarderPartie(*partie); err != nil {
		return err
	}
	return ajouterHistorique(*partie, resultat)
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
	for _, jouee := range partie.LettresJouees {
		runes := []rune(jouee)
		if len(runes) == 1 && unicode.ToLower(runes[0]) == lettre {
			return true
		}
	}
	return false
}
