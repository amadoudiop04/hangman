package hangman

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

// RechercheMot est la fonction principale implémentant la logique du jeu du Pendu.
func RechercheMot() {
    // Fichier contenant les mots
    filename := "../files/words.txt"

    // Charger les mots depuis le fichier
    words, err := ChargeDuMot(filename)
    if err != nil {
        fmt.Println("Erreur lors du chargement des mots:", err)
        return
    }

    rand.Seed(time.Now().UnixNano())

    // Choisir un mot au hasard parmi ceux chargés
    word := ChooseRandomWord(words)

    // Nombre de tentatives autorisées
    attempts := 10

    // Garder une trace des lettres révélées
    revealedLetters := make([]bool, len(word))

    // Afficher la position initiale du pendu
    DisplayPosition(10 - attempts)

    // Boucle principale du jeu
    for attempts > 0 {
        // Afficher l'état actuel du mot avec les lettres révélées
        DisplayWord(word, revealedLetters)
        fmt.Print("Suggérer une lettre: ")
        guess := ReadGuess()

        // Vérifier si la lettre proposée est dans le mot
        if strings.Contains(word, guess) {
            RevealLetters(word, guess, revealedLetters)
        } else {
            // Si la lettre proposée n'est pas dans le mot, réduire le nombre de tentatives et afficher la position du pendu
            attempts--
            fmt.Printf("La lettre '%s' n'est pas dans le mot. Tentatives restantes : %d\n", guess, attempts)
            DisplayPosition(10 - attempts)
        }

        // Vérifier si le mot a été entièrement deviné
        if IsWordComplete(revealedLetters) {
            fmt.Println("Félicitations! Vous avez trouvé le mot:", word)
            break
        }
    }

    // Si aucune tentative restante, afficher le mot correct
    if attempts == 0 {
        fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était:", word)
    }
}
