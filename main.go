package main

import (
	"fmt"
	"hangman/dictionnary"
	"hangman/hangman"
	"os"
)

func main() {

	err := dictionnary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load words: %v \n", err)
		os.Exit(1)
	}

	g := hangman.New(8, dictionnary.PickWord())

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
