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
	LettresJouees   []string `json:"lettres_jouees"`
	EssaiesRestants int      `json:"essaies_restants"`
}

func (p *Partie) selectionnerMot() error {
	data, err := os.ReadFile("mot.txt")
	if err != nil {
		return fmt.Errorf("impossible de lire mot.txt : %w", err)
	}
	var mots []string
	for _, mot := range strings.Fields(string(data)) {
		if WordIsValid(mot) {
			mots = append(mots, strings.ToLower(mot))
		}
	}
	if len(mots) == 0 {
		return fmt.Errorf("mot.txt ne contient aucun mot valide")
	}
	p.MotSecret = mots[rand.Intn(len(mots))]
	return nil
}
