package hangman

import (
	"fmt"
	"strings"
)

// AdvancedFeatures est une fonction qui effectue une boucle for simple et retourne la somme de deux entiers.
func AdvancedFeatures(a, b int) int {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	return a + b
}

// MotTrouve vérifie si le mot donné est contenu dans la phrase donnée et retourne un booléen.
func MotTrouve(phrase, mot string) bool {
	return strings.Contains(phrase, mot)
}

// Suite est une fonction de démonstration qui utilise MotTrouve pour rechercher un mot dans une phrase.
// Elle gère également un compteur de tentatives.
func Suite() {
	var tentatives int

	phrase := "phrase test"
	mot := " test "

	// Vérifie si le mot est trouvé dans la phrase.
	trouve := MotTrouve(phrase, mot)
	if trouve {
		fmt.Printf("Le mot a été trouvé dans la phrase. Le jeu s'arrête: %s\n", mot)
	} else {
		tentatives++

		maxTentatives := 2

		fmt.Printf("Le mot n'a pas été trouvé dans la phrase. Tentatives restantes: %d\n", maxTentatives-tentatives)
	}
}
