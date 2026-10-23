package main

import "fmt"

func (partie Partie) afficherPendu() {
	dessins := [...]string{
		` +---+
 |   |
     |
     |
     |
     |
=======`,
		` +---+
 |   |
 O   |
     |
     |
     |
=======`,
		` +---+
 |   |
 O   |
 |   |
     |
     |
=======`,
		` +---+
 |   |
 O   |
/|   |
     |
     |
=======`,
		` +---+
 |   |
 O   |
/|\  |
     |
     |
=======`,
		` +---+
 |   |
 O   |
/|\  |
/    |
     |
=======`,
		` +---+
 |   |
 O   |
/|\  |
/ \  |
     |
=======`,
	}
	erreurs := 6 - partie.EssaiesRestants
	if erreurs < 0 {
		erreurs = 0
	} else if erreurs >= len(dessins) {
		erreurs = len(dessins) - 1
	}
	fmt.Printf("\n%s\n", dessins[erreurs])
}
