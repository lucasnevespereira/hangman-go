package hangman

import "strings"

// Game => Structre of the Game
type Game struct {
	State        string   // Game State
	Letters      []string // Letters in word to find
	FoundLetters []string // Letters found in word
	UsedLetters  []string // Letters used
	TurnsLeft    int      // remaining guesses
}

// New inits the Game
func New(turns int, word string) *Game {
	// Remove spaces from word
	letters := strings.Split(strings.ToUpper(word), "")
	// Hide word with "_"
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	// Declare Game
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return g
}

// MakeAGuess allows users to make a guess
func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	}
}

// RevealLetter reveals all letters in word
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}

	return false
}
