package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)

	if n <= 0 {
		return result
	}

	// first row
	result[0] = []int{1}
	if n == 1 {
		return result
	}

	// second row
	result[1] = []int{1, 1}

	// third row onwards
	for row := 2; row < n; row++ {
		newRow := make([]int, row+1)
		newRow[0] = 1
		newRow[row] = 1
		for col := 1; col < row; col++ {
			newRow[col] = result[row-1][col-1] + result[row-1][col]
		}
		result[row] = newRow
	}

	return result
}
