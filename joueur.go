package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func demanderNom() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Veuillez renseigner votre nom : ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Erreur lors de la lecture du nom :", err)
			}
			return ""
		}

		nom := []rune(scanner.Text())
		if len(nom) == 0 {
			fmt.Println("Nom invalide, veuillez réessayer.")
			continue
		}

		nomValide := true
		for _, lettre := range nom {
			if !unicode.IsLetter(lettre) {
				nomValide = false
				break
			}
		}

		if !nomValide {
			fmt.Println("Nom invalide, veuillez réessayer.")
			continue
		}

		nom[0] = unicode.ToUpper(nom[0])
		for i := 1; i < len(nom); i++ {
			nom[i] = unicode.ToLower(nom[i])
		}

		return string(nom)
	}
}
