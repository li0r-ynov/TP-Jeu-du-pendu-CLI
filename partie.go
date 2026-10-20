package main

import "fmt"

type Partie struct {
	Nom             string   `json:"nom"`
	MotSecret       string   `json:"mot"`
	LettresJouees   []string `json:"lettres_jouees"`
	EssaiesRestants int      `json:"essaies_restants"`
}

func (p *Partie) selectionnerMot() error {
	mot, err := demarrerMot()
	if err != nil {
		return fmt.Errorf("sélection du mot : %w", err)
	}
	p.MotSecret = mot
	return nil
}
