package diamond

import (
	"errors"
	"strings"
)

func createRow(rowLetter byte, totalWidth int) string {
	var leadingSpaces int
	var middleSpaces int
	if rowLetter == 'A' {
		leadingSpaces = (totalWidth - 1) / 2
		middleSpaces = 0
		return strings.Repeat(" ", leadingSpaces) + string(rowLetter) + strings.Repeat(" ", middleSpaces) + strings.Repeat(" ", leadingSpaces)
	} else {
		middleSpaces = (2 * position(rowLetter)) - 1
		leadingSpaces = (totalWidth - 2 - middleSpaces) / 2
	}
	return strings.Repeat(" ", leadingSpaces) + string(rowLetter) + strings.Repeat(" ", middleSpaces) + string(rowLetter) + strings.Repeat(" ", leadingSpaces)
}

func calculateDimensions(letter byte) (width int) {
	total_width := 2*position(letter) + 1
	return total_width
}
func position(char byte) int {
	return int(char) - int('A')
}

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("character must be between 'A' and 'Z'")
	}
	width := calculateDimensions(char)
	var rows []string
	for letter := byte('A'); letter <= char; letter++ {
		rows = append(rows, createRow(letter, width))
	}
	for letter := char - 1; letter >= 'A'; letter-- {
		rows = append(rows, createRow(letter, width))
	}

	return strings.Join(rows, "\n"), nil
}
