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
