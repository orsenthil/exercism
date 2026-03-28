package matrix

import (
	"errors"
	"strconv"
	"strings"
)

func IsLargestInSlice(slice []int, value int) bool {
	for _, v := range slice {
		if v > value {
			return false
		}
	}
	return true
}

func IsSmallestInSlice(slice []int, value int) bool {
	for _, v := range slice {
		if v < value {
			return false
		}
	}
	return true
}

// Define the Matrix and Pair types here.

type Matrix [][]int
type Pair struct {
	row    int
	column int
}

func New(s string) (*Matrix, error) {
	if len(s) == 0 {
		return nil, nil
	}

	rows := strings.Split(strings.TrimSpace(s), "\n")
	matrix := make(Matrix, len(rows))

	for i, row := range rows {
		nums := strings.Fields(row)
		if i > 0 && len(nums) != len(matrix[0]) {
			return nil, errors.New("invalid matrix")
		}
		matrix[i] = make([]int, len(nums))
		for j, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			matrix[i][j] = n
		}
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	var saddlePoints []Pair

	if m == nil {
		return saddlePoints
	}

	for i, row := range *m {
		for j, value := range row {
			if IsLargestInSlice(row, value) && IsSmallestInSlice((*m).Column(j), value) {
				saddlePoints = append(saddlePoints, Pair{i + 1, j + 1})
			}
		}
	}
	return saddlePoints
}

func (m *Matrix) Column(j int) []int {
	column := make([]int, len(*m))
	for i, row := range *m {
		column[i] = row[j]
	}
	return column
}
