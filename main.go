package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
<<<<<<< HEAD
	if len(os.Args) < 2 {
		fmt.Println("Utilisation : go run . jouer | go run . reprendre")
=======
	if len(os.Args) == 2 && os.Args[1] == "historique" {
		if err := afficherHistorique(); err != nil {
			fmt.Fprintln(os.Stderr, "Impossible de lire l'historique :", err)
		}
		return
	}
	if len(os.Args) != 2 || (os.Args[1] != "jouer" && os.Args[1] != "reprendre") {
		fmt.Println("Utilisation : go run . jouer | reprendre | historique")
>>>>>>> 7647d7b4b28c2c851fb1a39db68d5e2164f7f5d1
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
<<<<<<< HEAD

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
=======
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
>>>>>>> 7647d7b4b28c2c851fb1a39db68d5e2164f7f5d1
			}
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
	if err := sauvegarderPartie(partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
		return
	}
	if err := jouer(scanner, &partie); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
	}
}

<<<<<<< HEAD
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
=======
func jouer(scanner *bufio.Scanner, partie *Partie) error {
	partie.afficherPendu()
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
>>>>>>> 7647d7b4b28c2c851fb1a39db68d5e2164f7f5d1
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
			partie.afficherPendu()
		}
<<<<<<< HEAD

		if err := sauvegarderPartie(*partie); err != nil {
			fmt.Fprintln(os.Stderr, "Erreur de sauvegarde :", err)
			return
=======
		if err := sauvegarderPartie(*partie); err != nil {
			return err
>>>>>>> 7647d7b4b28c2c851fb1a39db68d5e2164f7f5d1
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
