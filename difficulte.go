package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func demanderDifficulte(scanner *bufio.Scanner) string {
	for {
		fmt.Print("Choisissez la difficulté : 1 - Facile (jusqu'à 7 lettres), 2 - Difficile (8 lettres ou plus) : ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Erreur lors de la lecture de la difficulté :", err)
			}
			return ""
		}
		switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
		case "1", "facile":
			return "facile"
		case "2", "difficile":
			return "difficile"
		default:
			fmt.Println("Choix invalide. Entrez 1 ou 2.")
		}
	}
}

func difficulteMot(mot string) string {
	lettres := 0
	for _, lettre := range mot {
		if unicode.IsLetter(lettre) {
			lettres++
		}
	}
	if lettres <= 7 {
		return "facile"
	}
	return "difficile"
}
