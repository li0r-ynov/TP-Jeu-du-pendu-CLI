package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func demarrerMot() (string, error) {
	fichier, err := os.Open("mot.txt")
	if err != nil {
		return "", fmt.Errorf("impossible d'ouvrir mot.txt : %w", err)
	}
	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mot := strings.TrimSpace(scanner.Text())
		if WordIsValid(mot) {
			mots = append(mots, strings.ToLower(mot))
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("impossible de lire mot.txt : %w", err)
	}
	if len(mots) == 0 {
		return "", fmt.Errorf("mot.txt ne contient aucun mot")
	}

	mot := mots[rand.Intn(len(mots))]
	lettres := []rune(mot)
	var positions []int
	for i, lettre := range lettres {
		if unicode.IsLetter(lettre) {
			positions = append(positions, i)
		}
	}
	if len(positions) == 0 {
		return "", fmt.Errorf("le mot choisi ne contient aucune lettre")
	}

	lettreRevelee := unicode.ToLower(lettres[positions[rand.Intn(len(positions))]])
	masque := make([]string, len(lettres))
	for i, lettre := range lettres {
		if !unicode.IsLetter(lettre) || unicode.ToLower(lettre) == lettreRevelee {
			masque[i] = string(lettre)
		} else {
			masque[i] = "_"
		}
	}

	fmt.Printf("Mot à deviner : %s\n", strings.Join(masque, " "))
	return mot, nil
}

func WordIsValid(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter) && letter != '-' {
			return false
		}
	}
	return word != "" && !strings.HasPrefix(word, "-") && !strings.HasSuffix(word, "-") && !strings.Contains(word, "--")
}
