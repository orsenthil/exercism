package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows [][]int
}

func New(s string) (*Matrix, error) {
	matrix := &Matrix{}
	var length int
	for _, x := range strings.Split(s, "\n") {
		var row []int

		for _, cell := range strings.Split(strings.Trim(x, " "), " ") {
			n, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}

			row = append(row, int(n))
		}

		if length != 0 && len(row) != length {
			return nil, errors.New("length is not same")
		}
		length = len(row)

		matrix.rows = append(matrix.rows, row)

	}
	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {

	var matrix [][]int
	for i := 0; i < len(m.rows[0]); i++ {
		var col []int
		for j := 0; j < len(m.rows); j++ {
			col = append(col, m.rows[j][i])
		}

		matrix = append(matrix, col)
	}
	return matrix
}

func (m *Matrix) Rows() [][]int {
	var matrix [][]int
	for i := 0; i < len(m.rows); i++ {
		var row []int
		for j := 0; j < len(m.rows[0]); j++ {
			row = append(row, m.rows[i][j])
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= 0 && col >= 0 && row < len(m.rows) && col < len(m.rows[0]) {
		m.rows[row][col] = val
		return true
	}
	return false
}
