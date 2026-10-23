package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Utilisation : go run . jouer | go run . reprendre")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	var partie Partie
	switch os.Args[1] {
	case "jouer":
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
	case "reprendre":
		var err error
		partie, err = chargerPartie()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur :", err)
			return
		}
		fmt.Printf("Reprise de la partie de %s.\n", partie.Nom)
	default:
		fmt.Println("Utilisation : go run . jouer | go run . reprendre")
		return
	}

	if err := sauvegarderPartie(partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
		return
	}

	jouerPartie(scanner, &partie)
}

func jouerPartie(scanner *bufio.Scanner, partie *Partie) {
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
		quitter, err := gererCommandeSauvegarde(*partie, saisie)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
			return
		}
		if quitter {
			fmt.Println("Partie sauvegardée. À bientôt !")
			return
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

		if err := sauvegarderPartie(*partie); err != nil {
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
