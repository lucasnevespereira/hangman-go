package dictionnary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

// Load is a func to load words from a filename
func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// PickWord picks a word randomnly
func PickWord() string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
