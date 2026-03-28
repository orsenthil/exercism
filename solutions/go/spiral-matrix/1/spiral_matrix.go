package spiralmatrix

func SpiralMatrix(size int) [][]int {
	// Handle edge case.
	if size == 0 {
		return [][]int{}
	}

	matrix := make([][]int, size)

	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// Moving right example
	row := 0
	col := 0

	current := 1

	top := 0
	bottom := size - 1
	left := 0
	right := size - 1

	for current <= size*size {
		// Move right
		for col <= right && current <= size*size {
			matrix[row][col] = current
			current++
			col++
		}
		col-- // Fix position
		top++ // Update boundary
		row++ // Prepare for next direction

		// Move down
		for row <= bottom && current <= size*size {
			matrix[row][col] = current
			current++
			row++
		}
		row--   // Fix position
		right-- // Update boundary
		col--   // Prepare for next direction

		// Move left
		for col >= left && current <= size*size {
			matrix[row][col] = current
			current++
			col--
		}
		col++    // Fix position
		bottom-- // Update boundary
		row--    // Prepare for next direction

		// Move up
		for row >= top && current <= size*size {
			matrix[row][col] = current
			current++
			row--
		}
		row++  // Fix position
		left++ // Update boundary
		col++  // Prepare for next direction
	}

	return matrix

}
