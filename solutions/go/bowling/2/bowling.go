// Package bowling provides functionality for calcualting bowling scores.

package bowling

import "errors"

// Game represents a bowling game.

type Game struct {
	rolls []rollResult
}

type rollResult struct {
	pins int

	frame int

	throw int

	isStrike bool

	isSpare bool
}

// NewGame generates an instance of the game struct and returns the pointer to that.

func NewGame() *Game {

	return &Game{rolls: make([]rollResult, 0)}

}

// Roll adds a roll result to the underlying game.

func (g *Game) Roll(pins int) error {

	if pins < 0 || 10 < pins {

		return errors.New("pins must be between 0 and 10")

	}

	if len(g.rolls) == 0 {

		g.rolls = append(g.rolls, rollResult{pins: pins, frame: 1, throw: 1, isStrike: pins == 10})

		return nil

	}

	lastRoll := g.rolls[len(g.rolls)-1]

	if lastRoll.throw == 1 && lastRoll.pins != 10 && lastRoll.pins+pins > 10 {

		return errors.New("the sum of first and second pins cannot exceed 10")

	}

	if lastRoll.frame < 10 {

		if lastRoll.throw == 1 && lastRoll.isStrike {

			g.rolls = append(g.rolls, rollResult{pins: pins, frame: lastRoll.frame + 1, throw: 1, isStrike: pins == 10 && lastRoll.frame != 9})

		} else if lastRoll.throw == 1 {

			g.rolls = append(g.rolls, rollResult{pins: pins, frame: lastRoll.frame, throw: 2, isSpare: lastRoll.pins+pins == 10})

		} else {

			g.rolls = append(g.rolls, rollResult{pins: pins, frame: lastRoll.frame + 1, throw: 1, isStrike: pins == 10 && lastRoll.frame != 9})

		}

	} else if lastRoll.throw == 3 {

		return errors.New("game is already over")

	} else if lastRoll.throw == 2 {

		if firstThrow := g.rolls[len(g.rolls)-2]; firstThrow.pins+lastRoll.pins < 10 {

			return errors.New("game is already over")

		} else if firstThrow.pins == 10 && lastRoll.pins != 10 && lastRoll.pins+pins > 10 {

			return errors.New("pin count exceeds 10")

		}

		g.rolls = append(g.rolls, rollResult{pins: pins, frame: 10, throw: 3})

	} else {

		g.rolls = append(g.rolls, rollResult{pins: pins, frame: 10, throw: 2})

	}

	return nil

}

// Score calculates the score of a bowling game.

func (g *Game) Score() (int, error) {

	var result int

	if len(g.rolls) == 0 {

		return 0, errors.New("Score cannot be taken until the end of the game")

	}

	lastRoll := g.rolls[len(g.rolls)-1]

	if lastRoll.frame != 10 || (lastRoll.frame == 10 && lastRoll.throw < 2) {

		return 0, errors.New("Score cannot be taken until the end of the game")

	} else if lastRoll.frame == 10 && lastRoll.throw == 2 {

		if firstThrow := g.rolls[len(g.rolls)-2]; firstThrow.pins+lastRoll.pins >= 10 {

			return 0, errors.New("Score cannot be taken until the end of the game")

		}

	}

	for i := 0; i < len(g.rolls); i++ {

		if g.rolls[i].isStrike {

			result += g.rolls[i].pins + g.rolls[i+1].pins + g.rolls[i+2].pins

		} else if g.rolls[i].isSpare {

			result += g.rolls[i].pins + g.rolls[i+1].pins

		} else {

			result += g.rolls[i].pins

		}

	}

	return result, nil

}
