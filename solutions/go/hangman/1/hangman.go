package hangman

import (
	"errors"
	"slices"
)

type Game struct {
	word              string
	guesses           []rune
	remainingFailures int
}

func NewGame(word string) *Game {
	return &Game{
		word:              word,
		guesses:           make([]rune, 0, 10),
		remainingFailures: 9,
	}
}

func (g *Game) Guess(r rune) error {
	if g.State() == "Lose" {
		return errors.New("cannot guess after the game is lost")
	}
	if g.State() == "Win" {
		return errors.New("cannot guess after the game is won")
	}

	if slices.Contains(g.guesses, r) {
		g.remainingFailures--
		return nil
	}
	for _, letter := range g.word {
		if letter == r {
			g.guesses = append(g.guesses, r)
			return nil
		}
	}
	g.remainingFailures--
	return nil
}

func (g *Game) MaskedWord() string {
	result := []rune{}
	for _, letter := range g.word {
		if slices.Contains(g.guesses, letter) {
			result = append(result, letter)
		} else {
			result = append(result, '_')
		}
	}
	return string(result)
}

func (g *Game) RemainingGuesses() int {
	if g.remainingFailures < 0 {
		return 0
	}
	return g.remainingFailures
}

func (g *Game) State() string {
	if g.word == g.MaskedWord() {
		return "Win"
	}
	if g.remainingFailures < 0 {
		return "Lose"
	}
	return "Ongoing"
}
