# Jeu du pendu en CLI

Un jeu du pendu développé en Go, jouable dans le terminal. Le but est de deviner un mot lettre par lettre avant de faire 6 erreurs.

## Fonctionnalités

- Deux difficultés : facile (jusqu’à 7 lettres) et difficile (8 lettres ou plus).
- Affichage du pendu dans le terminal.
- Sauvegarde et reprise d’une partie.
- Historique des parties terminées.

## Utilisation

Installez Go 1.25 ou une version plus récente, puis ouvrez un terminal dans le dossier du projet.

```sh
go run . jouer
```

Choisissez un pseudo et une difficulté, puis proposez une lettre à chaque tour. Tapez `!save` pour sauvegarder et quitter.

Pour reprendre la partie sauvegardée :

```sh
go run . reprendre
```

Pour consulter l’historique :

```sh
go run . historique
```

## Fichiers du projet

- `main.go` : lance le programme et gère le déroulement du jeu.
- `joueur.go` : demande et vérifie le nom du joueur.
- `difficulte.go` : gère le choix de la difficulté selon la longueur des mots.
- `startmot.go` : lit la liste des mots et en choisit un au hasard.
- `partie.go` : contient les informations de la partie et révèle une lettre au départ.
- `pendu.go` : affiche le dessin du pendu selon le nombre d’erreurs.
- `save.go` : sauvegarde et recharge une partie.
- `sauvegarde.go` : enregistre et affiche l’historique des parties.
- `mot.txt` : contient les mots à deviner.
- `save.json` : contient la partie sauvegardée.
- `score.csv` : contient l’historique des résultats.
- `go.mod` : définit le module du projet et la version de Go.
- `README.md` : présente le projet et son utilisation.
