package main

import (
	"fmt"
	"math/rand"
	"unicode"
)

type Partie struct {
	Nom             string   `json:"nom"`
	Difficulte      string   `json:"difficulte"`
	MotSecret       string   `json:"mot"`
	LettresJouees   []string `json:"lettres_jouees"`
	EssaiesRestants int      `json:"essaies_restants"`
}

func (p *Partie) selectionnerMot() error {
	mot, err := demarrerMot(p.Difficulte)
	if err != nil {
		return fmt.Errorf("sélection du mot : %w", err)
	}
	p.MotSecret = mot

	runes := []rune(mot)
	var positions []int
	for i, lettre := range runes {
		if unicode.IsLetter(lettre) {
			positions = append(positions, i)
		}
	}
	if len(positions) == 0 {
		return fmt.Errorf("le mot choisi ne contient aucune lettre")
	}
	position := positions[rand.Intn(len(positions))]
	lettreRevelee := unicode.ToLower(runes[position])
	p.LettresJouees = append(p.LettresJouees, string(lettreRevelee))
	return nil
}
