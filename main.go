package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		afficherUtilisation()
		return
	}

	if os.Args[1] == "historique" {
		if err := afficherHistorique(); err != nil {
			fmt.Fprintln(os.Stderr, "Impossible de lire l'historique :", err)
		}
		return
	}
	if os.Args[1] != "jouer" && os.Args[1] != "reprendre" {
		afficherUtilisation()
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
		if partie.EssaiesRestants == 0 || partie.motTrouve() {
			if partie.EssaiesRestants == 0 {
				fmt.Printf("Cette partie est terminée. Le mot était : %s\n", partie.MotSecret)
			} else {
				fmt.Printf("Cette partie est déjà gagnée : %s\n", partie.MotSecret)
			}
			return
		}
		fmt.Printf("Reprise de la partie de %s.\n", partie.Nom)
	} else {
		nom := demanderNom(scanner)
		if nom == "" {
			return
		}
		partie = Partie{
			Nom:             nom,
			LettresJouees:   []string{},
			EssaiesRestants: 6,
		}
		if err := partie.selectionnerMot(); err != nil {
			fmt.Fprintln(os.Stderr, "Erreur :", err)
			return
		}
		fmt.Printf("Bonjour %s !\n", nom)
	}

	if err := sauvegarderPartie(partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
		return
	}
	if err := jouer(scanner, &partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}

func afficherUtilisation() {
	fmt.Println("Utilisation : go run . jouer | go run . reprendre | go run . historique")
}

func jouer(scanner *bufio.Scanner, partie *Partie) error {
	partie.afficherPendu()
	for partie.EssaiesRestants > 0 && !partie.motTrouve() {
		fmt.Printf("\nMot à deviner : %s\nEssais restants : %d\nEntrez une lettre (!save pour sauvegarder et quitter) : ",
			partie.afficherMot(), partie.EssaiesRestants)

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("lecture de la lettre : %w", err)
			}
			fmt.Println("\nPartie interrompue et sauvegardée. Relancez avec `go run . reprendre`.")
			return sauvegarderPartie(*partie)
		}

		saisie := strings.TrimSpace(scanner.Text())
		quitter, err := gererCommandeSauvegarde(*partie, saisie)
		if err != nil {
			return err
		}
		if quitter {
			fmt.Println("Partie sauvegardée. À bientôt !")
			return nil
		}

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
		presente := false
		for _, lettreMot := range partie.MotSecret {
			if unicode.ToLower(lettreMot) == lettre {
				presente = true
				break
			}
		}
		if !presente {
			partie.EssaiesRestants--
			fmt.Println("Cette lettre n'est pas dans le mot.")
			partie.afficherPendu()
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
