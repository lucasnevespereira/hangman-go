package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(l)
}
