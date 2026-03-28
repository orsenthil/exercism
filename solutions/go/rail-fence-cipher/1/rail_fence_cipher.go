package railfence

func createRailFence(message string, rails int) [][]rune {
	// Create a 2D slice to hold the rails
	fence := make([][]rune, rails)

	for i := range fence {
		fence[i] = make([]rune, len(message))
		for j := range fence[i] {
			fence[i][j] = '.' // Initialize the fence with dots
		}
	}
	return fence
}

func zigzagTraversal(length, rails int) []int {
	positions := make([]int, length)
	rail := 0
	goingDown := true

	for i := 0; i < length; i++ {
		positions[i] = rail
		if goingDown {
			rail++
			if rail == rails-1 {
				goingDown = false
			}
		} else {
			rail--
			if rail == 0 {
				goingDown = true
			}
		}
	}
	return positions
}

func Encode(message string, rails int) string {
	fence := createRailFence(message, rails)
	zigzag := zigzagTraversal(len(message), rails)

	for i, char := range message {
		fence[zigzag[i]][i] = char
	}

	encoded := ""
	for _, rail := range fence {
		for _, char := range rail {
			if char != '.' {
				encoded += string(char)
			}
		}
	}
	return encoded
}

func Decode(message string, rails int) string {
	fence := createRailFence(message, rails)
	zigzag := zigzagTraversal(len(message), rails)

	// Step 1: Mark the positions of the characters in the fence
	for i := 0; i < len(message); i++ {
		fence[zigzag[i]][i] = '?'
	}

	// Step 2: Fill the marked positions row by row
	index := 0
	for i := range fence {
		for j := range fence[i] {
			if fence[i][j] == '?' {
				fence[i][j] = rune(message[index])
				index++
			}
		}
	}
	// Step 3 : Traverse the fence in zigzag order to get the decoded message
	result := make([]rune, len(message))
	for i := 0; i < len(message); i++ {
		result[i] = fence[zigzag[i]][i]
	}

	return string(result)
}
