package gameoflife

func Tick(matrix [][]int) [][]int {
	result := make([][]int, len(matrix))
	for i := range matrix {
		result[i] = make([]int, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			result[i][j] = Conway(matrix[i][j], i, j, matrix)
		}
	}
	return result
}

func Conway(cell int, i int, j int, matrix [][]int) int {
	neighbors := 0
	direction := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range direction {
		if i+dir[0] >= 0 && i+dir[0] < len(matrix) && j+dir[1] >= 0 && j+dir[1] < len(matrix[i]) && matrix[i+dir[0]][j+dir[1]] == 1 {
			neighbors++
		}
	}
	if cell == 1 {
		if neighbors == 2 || neighbors == 3 {
			return 1
		}
		return 0
	} else {
		if neighbors == 3 {
			return 1
		}
		return 0
	}
}
