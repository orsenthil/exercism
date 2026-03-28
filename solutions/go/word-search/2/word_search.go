package wordsearch

import "errors"

func checkWord(puzzle []string, word string, row, col int, rowDirection, colDirection int) bool {
	if len(word) == 0 {
		return false
	}
	endRow := row + (len(word)-1)*rowDirection
	endCol := col + (len(word)-1)*colDirection

	if endRow < 0 || endRow >= len(puzzle) || endCol < 0 || endCol >= len(puzzle[0]) {
		return false
	}
	for i := 0; i < len(word); i++ {
		currentRow := row + i*rowDirection
		currentCol := col + i*colDirection
		if puzzle[currentRow][currentCol] != word[i] {
			return false
		}
	}
	return true
}

func getCoordinates(row, col, rowDirection, colDirection int, wordLen int) [2][2]int {
	endRow := row + (wordLen-1)*rowDirection
	endCol := col + (wordLen-1)*colDirection
	coords := [2][2]int{{col, row}, {endCol, endRow}}
	return coords
}

func findWord(word string, puzzle []string, directions [][2]int) ([2][2]int, bool) {
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0]); col++ {
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if checkWord(puzzle, word, row, col, dx, dy) {
					return getCoordinates(row, col, dx, dy, len(word)), true
				}
			}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}, false
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	hasError := false

	directions := [][2]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // diagonal down-right
		{-1, -1}, // diagonal up-left
		{1, -1},  // diagonal down-left
		{-1, 1},  // diagonal up-right
	}

	for _, word := range words {
		coords, found := findWord(word, puzzle, directions)
		result[word] = coords
		if !found {
			hasError = true
		}
	}
	if hasError {
		return result, errors.New("word not found")
	}
	return result, nil
}
