package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ChargeDuMot lit un fichier et renvoie une tranche de chaînes contenant les mots.
func ChargeDuMot(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// chooseRandomWord sélectionne un mot aléatoire dans une tranche de mots.
func ChooseRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

// displayWord affiche l'état actuel du mot avec les lettres révélées et non révélées.
func DisplayWord(word string, revealed []bool) {
	fmt.Print("Mot : ")
	for i, char := range word {
		if revealed[i] {
			fmt.Printf("%c ", char)
		} else {	
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

// revealLetters met à jour la tranche révélée en fonction de la lettre devinée.
func RevealLetters(word, guess string, revealed []bool) {
	for i, char := range word {
		if string(char) == guess {
			revealed[i] = true
		}
	}
}

// isWordComplete vérifie si toutes les lettres du mot sont révélées.
func IsWordComplete(revealed []bool) bool {
	for _, val := range revealed {
		if !val {
			return false
		}
	}
	return true
}

// readGuess lit une supposition de l'utilisateur depuis la console.
func ReadGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')
	return strings.TrimSpace(guess)
}
