package transpose

func maxLength(lines []string) int {
	maxl := 0
	for _, line := range lines {
		if len(line) > maxl {
			maxl = len(line)
		}
	}
	return maxl
}

func getChar(lines string, i int) byte {
	if i >= len(lines) {
		return ' '
	}
	return lines[i]
}

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	numOutputRows := maxLength(input)
	numOutputCols := len(input)
	result := make([]string, numOutputRows)

	maxNeededLength := make([]int, numOutputRows)
	for row := 0; row < numOutputRows; row++ {
		for col := 0; col < numOutputCols; col++ {
			if row < len(input[col]) {
				maxNeededLength[row] = col + 1
			}
		}
	}

	for row := 0; row < numOutputRows; row++ {
		currentRow := make([]byte, maxNeededLength[row])
		for col := 0; col < maxNeededLength[row]; col++ {
			currentRow[col] = getChar(input[col], row)
		}
		result[row] = string(currentRow)
	}

	return result
}
