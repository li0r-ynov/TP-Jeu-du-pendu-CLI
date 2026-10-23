package main

import "fmt"

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
	return nil
}
