package flowerfield

func Annotate(board []string) []string {
	if len(board) == 0 {
		return board
	}

	result := make([]string, len(board))

	for r, row := range board {
		buf := []byte(row)
		for c := range buf {
			if buf[c] == '*' {
				continue
			} else {
				count := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr, nc := r + dr , c + dc
						if nr >= 0 && nc >= 0 && nr < len(board) && nc < len(board[nr]) {
							if board[nr][nc] == '*' {
								count += 1
							}
						}
					}
				}
				if count > 0 {
					buf[c] = byte('0' + count)
				}
			}
		}
		result[r] = string(buf)
	}

	return result
}
