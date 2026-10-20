package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Partie struct {
	Nom             string   `json:"nom"`
	MotSecret       string   `json:"mot"`
	LettresJouees   []string `json:"lettres_jouées"`
	EssaiesRestants int      `json:"essaies_restants"`
}

func (g *Partie) selectionnerMot() error {
	data, err := os.ReadFile("mot.txt")
	if err != nil {
		return fmt.Errorf("impossible de lire mot.txt : %w", err)
	}

	mots := strings.Fields(string(data))
	if len(mots) == 0 {
		return fmt.Errorf("mot.txt ne contient aucun mot")
	}

	g.MotSecret = mots[rand.Intn(len(mots))]
	return nil
}
