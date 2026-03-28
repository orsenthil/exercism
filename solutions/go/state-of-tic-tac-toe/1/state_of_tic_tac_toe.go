package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func countPieces(board []string) (xs, os int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == 'X' {
				xs++
			} else if cell == 'O' {
				os++
			}
		}
	}
	return
}

func isWinningRow(row string) bool {
	return row == "XXX" || row == "OOO"
}

func hasWin(board []string) (bool, rune) {
	// Check rows
	for _, row := range board {
		if row == "XXX" {
			return true, 'X'
		}
		if row == "OOO" {
			return true, 'O'
		}
	}
	// TODO: Add column and diagonal checks
	return false, ' '
}

func StateOfTicTacToe(board []string) (State, error) {
	// 1. First count X's and O's
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case 'X':
				xCount++
			case 'O':
				oCount++
			case ' ':
				// valid empty space
			default:
				// handle invalid characters if needed
			}
		}
	}

	// 2. Validate the counts
	if oCount > xCount || xCount > oCount+1 {
		return "", fmt.Errorf("invalid game: wrong number of moves")
	}
}
