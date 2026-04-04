package killersudokuhelper

import "slices"

func Combinations(sum, size int, exclude []int) [][]int {
	var result [][]int
	var find func(start, size, sum int, current[]int)

	find = func(start, size, sum int, current []int) {
		if size == 0 {
			if sum == 0 {
				result = append(result, append([]int{}, current...))
			}
			return
		}
		for digit := start; digit <= 9; digit++ {
			if slices.Contains(exclude, digit) {
				continue
			}
			find(digit+1, size-1, sum-digit, append(current, digit))
		}
	}
	find(1, size, sum, nil)
	return result
}
