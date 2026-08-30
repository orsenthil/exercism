package savethecow

import (
	"errors"
	"strings"
)

type Game struct{
	word string
	guessed map[rune]bool
	remainingGuesses int
	lost bool
}

func NewGame(word string) *Game {
	return &Game{word: word, guessed: make(map[rune]bool), remainingGuesses: 9}
}

func (g *Game) Guess(r rune) error {
	if g.State() == "Win" {
		return errors.New("cannot guess after the game is won")
	}
	if g.State() == "Lose" {
		return errors.New("cannot guess after the game is lost")
	}
	repeated := g.guessed[r]
	g.guessed[r] = true
	if repeated || !strings.ContainsRune(g.word, r) {
		if g.remainingGuesses == 0 {
			g.lost = true
			return nil
		} else {
			g.remainingGuesses -= 1
		}
	}
	return nil
}

func (g *Game) MaskedWord() string {
	var out []rune
	for _, ch := range g.word {
		if g.guessed[ch] {
			out = append(out, ch)
		} else {
			out = append(out, '_')
		}
	}
	return string(out)
}

func (g *Game) RemainingGuesses() int {
	return g.remainingGuesses
}

func (g *Game) fullyRevealed() bool {
	for _, ch := range g.word {
		if !g.guessed[ch] {
			return false
		}
	}
	return true
}

func (g *Game) State() string {
	if g.fullyRevealed() {
		return "Win"
	}
	if g.lost {
		return "Lose"
	}
	return "Ongoing"
}
