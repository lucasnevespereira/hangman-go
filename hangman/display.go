package hangman

import (
	"fmt"

	"github.com/fatih/color"
)

// DrawWelcome draws welcoming game message
func DrawWelcome() {
	color.Cyan("HANGMANG-GO")
}

// Draw current state of the game
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
	 / \  |
				|
	=========
		`

	case 1:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
	 /    |
				|
	=========
		`
	case 2:
		draw = `
		+---+
		|   |
		O   |
	 /|\  |
				|
				|
	=========
		`
	case 3:
		draw = `
		+---+
		|   |
		O   |
	 /|   |
				|
				|
	=========
		`
	case 4:
		draw = `
		+---+
		|   |
		O   |
		|   |
				|
				|
	=========
		`
	case 5:
		draw = `
		+---+
		|   |
		O   |
				|
				|
				|
	=========
		`
	case 6:
		draw = `
		+---+
		|   |
				|
				|
				|
				|
	=========
		`
	case 7:
		draw = `
			  |
				|
				|
				|
				|
	========='
		`

	case 8:
		draw = `

		`
	}

	color.Red(draw)
}

func drawState(g *Game, guess string) {
	color.Green("Guessed:")
	drawLetters(g.FoundLetters)
	fmt.Println()
	fmt.Println()

	color.Yellow("Used:")
	drawLetters(g.UsedLetters)
	fmt.Println()
	fmt.Println()

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess,  '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost 🙁 . The word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON 😀 . The word was:")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
}
