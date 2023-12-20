package main

import (
	"hangman-classic"
)

func main() {
    hangman.RechercheMot()
	hangman.DisplayPosition(10)
	hangman.AdvancedFeatures(3, 5) 
	game := "hangmanGame"
	file := "../Ascii letters/standard.txt"
	hangman.AsciiArt(file, game)
	hangman.StartAndStop()
}