package minesweeper

func getAdjacentPositions(row, col int) [][2]int {
	return [][2]int{
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col - 1}, {row, col + 1},
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}
}

func isValidPosition(row, col int, board []string) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return board
	}
	rows, cols := len(board), len(board[0])

	annotatedBoard := make([]string, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '*' {
				annotatedBoard[i] += "*"
				continue
			}

			count := 0
			for _, pos := range getAdjacentPositions(i, j) {
				if isValidPosition(pos[0], pos[1], board) && board[pos[0]][pos[1]] == '*' {
					count++
				}
			}

			if count == 0 {
				annotatedBoard[i] += " "
			} else {
				annotatedBoard[i] += string(rune('0' + count))
			}
		}
	}
	return annotatedBoard
}
