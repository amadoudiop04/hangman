package hangman

import (
	"bufio"
	"fmt"
	"os"
)

type HangmanGame struct {
	// Définissez ici les champs de votre structure hangmanGame
	WordDisplayed string
	// ... autres champs
}

func AsciiArt(file string, word string) {
	// Ouverture du fichier
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer readFile.Close()

	// Création du scanner pour lire le fichier ligne par ligne
	fileScanner := bufio.NewScanner(readFile)

	var slice []string

	// Lecture du fichier ligne par ligne
	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}

	// Affichage du mot en ASCII ligne par ligne
	runeWord := []rune(word)
	for k := 0; k < 7; k++ {
		for i := 0; i < len(runeWord); i++ {
			if 97 <= runeWord[i] && runeWord[i] <= 122 {
				fmt.Print(slice[int(runeWord[i]-97)*7+k])
			} else if 65 <= runeWord[i] && runeWord[i] <= 90 {
				fmt.Print(slice[int(runeWord[i]-65)*7+k])
			} else {
				fmt.Print(slice[182+k])
			}
		}
		fmt.Println("")
	}
}
