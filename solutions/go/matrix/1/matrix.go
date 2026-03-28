package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	values [][]int
}

func convertToInt(strcols []string) ([]int, error) {
	var res []int
	for _, s := range strcols {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	var cols []int
	var err error

	var mat Matrix

	for _, row := range rows {
		strcols := strings.Split(row, " ")
		cols, err = convertToInt(strcols)
		if err != nil {
			return nil, errors.New("Matrix Creation Error.")
		}
		mat.values = append(mat.values, cols)
	}

	return &mat, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	rows := len(m.values)
	cols := len(m.values[0])
	mat := make([][]int, cols)

	for i := 0; i < cols; i++ {
		mat[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			mat[j][i] = m.values[i][j]
		}
	}
	return mat
}

func (m *Matrix) Rows() [][]int {
	rows := len(m.values)
	cols := len(m.values[0])

	mat := make([][]int, rows)

	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			mat[i][j] = m.values[i][j]
		}
	}

	return mat
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= len(m.values) || col >= len(m.values[0]) {
		return false
	}
	m.values[row][col] = val
	return true
}
