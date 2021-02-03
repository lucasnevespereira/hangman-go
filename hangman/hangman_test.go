package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"s", "a", "m"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got %v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"s", "a", "m"}
	guess := "o"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s. got %v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(2, "")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "sam")
	g.MakeAGuess("a")
	validState(t, "goodGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "sam")
	g.MakeAGuess("a")
	g.MakeAGuess("a")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "sam")
	g.MakeAGuess("b")
	validState(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "sas")
	g.MakeAGuess("s")
	g.MakeAGuess("a")
	g.MakeAGuess("s")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "samantha")
	g.MakeAGuess("a")
	g.MakeAGuess("t")
	g.MakeAGuess("e")
	g.MakeAGuess("s")
	g.MakeAGuess("t")
	g.MakeAGuess("i")
	g.MakeAGuess("n")
	g.MakeAGuess("g")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. got=%v", expectedState, actualState)
		return false
	}

	return true
}
