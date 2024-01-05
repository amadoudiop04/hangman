package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GameState struct {
	Word     string
	Guesses  []string
	Attempts int
}

func StartAndStop() {
	var game GameState

	// Vérifiez si le jeu doit être démarré avec une sauvegarde précédente
	if len(os.Args) == 2 && os.Args[1] == "--startWith" {
		LoadGame(&game, "save.txt")
	} else {
		// Initialiser un nouveau jeu
		game = GameState{
			Word:     "hangman",
			Guesses:  make([]string, 0),
			Attempts: 6,
		}
	}

	// Boucle principale du jeu
	for {
		DisplayGameState(game)

		// Lire l'entrée utilisateur
		var input string
		fmt.Print(" 'STOP' pour quitter: ")
		fmt.Scanln(&input)

		// Vérifier si l'utilisateur veut arrêter le jeu
		if input == "STOP" {
			SaveGame(game, "save.txt")
			fmt.Println("Jeu enregistré. Au revoir!")
			os.Exit(0)
		}

		// Autres logiques de jeu ici (vérification des lettres, gestion des tentatives, etc.)
		// ...

		// Exemple: Ajouter la lettre à la liste des suppositions
		game.Guesses = append(game.Guesses, input)
	}
}

func DisplayGameState(game GameState) {
	// Afficher l'état actuel du jeu
	// ...
}

func SaveGame(game GameState, filename string) {
	// Convertir l'état du jeu en JSON
	data, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}

	// Écrire les données JSON dans un fichier
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier de sauvegarde:", err)
	}
}

func LoadGame(game *GameState, filename string) {
	// Lire les données JSON à partir du fichier
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de sauvegarde:", err)
		return
	}

	// Décoder les données JSON dans la structure du jeu
	err = json.Unmarshal(data, game)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
	}
}
