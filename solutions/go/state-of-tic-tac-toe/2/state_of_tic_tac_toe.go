package stateoftictactoe

import (
	"fmt"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func availableMove(board []string) bool {
	for _, line := range board {
		if strings.Contains(line, " ") {
			return true
		}
	}
	return false
}

func validBoard(board []string) bool {
	var countX, countO int
	for _, line := range board {
		countX += strings.Count(line, "X")
		countO += strings.Count(line, "O")
	}
	if countX == countO || countX == countO+1 {
		return true
	}
	return false
}

// StateOfTicTacToe returns the state of the game
func StateOfTicTacToe(board []string) (State, error) {

	// check the correct sequence of moves
	if !validBoard(board) {
		return "", fmt.Errorf("invalid sequence of moves")
	}

	var timesWinnerX, timesWinnerO bool

	// building the array with all the valid combinations
	trisCombs := make([][]int, 0, 8)
	trisCombs = append(trisCombs, []int{0, 0, 1, 1, 2, 2})
	trisCombs = append(trisCombs, []int{0, 2, 1, 1, 2, 0})
	for i := 0; i < 3; i++ {
		trisCombs = append(trisCombs, []int{i, 0, i, 1, i, 2})
		trisCombs = append(trisCombs, []int{0, i, 1, i, 2, i})
	}
	// checking all valid combinations
	for _, comb := range trisCombs {
		if board[comb[0]][comb[1]] != ' ' &&
			board[comb[0]][comb[1]] == board[comb[2]][comb[3]] &&
			board[comb[2]][comb[3]] == board[comb[4]][comb[5]] {
			if board[comb[0]][comb[1]] == 'X' {
				timesWinnerX = true
			} else {
				timesWinnerO = true
			}
		}
	}

	switch {
	case timesWinnerX && timesWinnerO:
		return "", fmt.Errorf("too many winners")
	case timesWinnerX || timesWinnerO:
		return Win, nil
	default:
		if availableMove(board) {
			return Ongoing, nil
		}
		return Draw, nil
	}
}
