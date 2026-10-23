package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func demarrerMot(difficulte string) (string, error) {
	if difficulte != "facile" && difficulte != "difficile" {
		return "", fmt.Errorf("difficulté invalide : %s", difficulte)
	}
	fichier, err := os.Open("mot.txt")
	if err != nil {
		return "", fmt.Errorf("impossible d'ouvrir mot.txt : %w", err)
	}
	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mot := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if WordIsValid(mot) && difficulteMot(mot) == difficulte {
			mots = append(mots, mot)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("impossible de lire mot.txt : %w", err)
	}
	if len(mots) == 0 {
		return "", fmt.Errorf("mot.txt ne contient aucun mot valide pour la difficulté %s", difficulte)
	}
	return mots[rand.Intn(len(mots))], nil
}

func WordIsValid(word string) bool {
	if word == "" || strings.HasPrefix(word, "-") || strings.HasSuffix(word, "-") || strings.Contains(word, "--") {
		return false
	}
	for _, letter := range word {
		if !unicode.IsLetter(letter) && letter != '-' {
			return false
		}
	}
	return true
}
