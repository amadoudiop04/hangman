package hangman

import (
	"bufio"
	"fmt"
	"os"
)

// DisplayPosition affiche l'état actuel du jeu du pendu en fonction du stage fourni.
func DisplayPosition(stage int) {
	// Spécifie le chemin vers le fichier hangman.txt
	hangmanFile := "../hangman.txt"

	// Tente d'ouvrir le fichier hangman.txt
	file, err := os.Open(hangmanFile)
	if err != nil {
		// Affiche un message d'erreur s'il y a un problème lors de l'ouverture du fichier
		fmt.Println("Erreur lors de l'ouverture du fichier hangman.txt :", err)
		return
	}
	// Ferme le fichier lorsque la fonction est terminée (reporté grâce à "defer")
	defer file.Close()

	// Crée un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	// Boucle pour sauter les lignes dans le fichier jusqu'à ce que le stage spécifié soit atteint
	for i := 0; i < stage; i++ {
		// Boucle interne pour sauter huit lignes pour chaque stage
		for j := 0; j < 8; j++ {
			// Lit et ignore la ligne actuelle
			scanner.Scan()
		}
	}

	// Boucle pour afficher les huit lignes suivantes, représentant l'état actuel du pendu
	for j := 0; j < 8; j++ {
		// Lit et affiche la ligne actuelle
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
}
