package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents integers in rows and columns
type Matrix [][]int

// New returns a matrix of integers
func New(input string) (Matrix, error) {
	// instantiate a new matrix 
	m := Matrix{}
	rowSize := -1

	for i, row := range strings.Split(input, "\n") {
		vs := strings.Fields(row)

		if rowSize == -1 {
			rowSize = len(vs)
		} else if rowSize != len(vs) {
			return nil, errors.New("error")
		}

		m = append(m, []int{})
		for _, v := range vs {
			if num, err := strconv.Atoi(v); err == nil {
				m[i] = append(m[i], num)
			} else {
				return nil, err
			}
		}
	}
	return m, nil
}

// Rows returns a list of matrix rows
func (m Matrix) Rows() [][]int {
	if len(m) == 0 {
		return [][]int{}
	}

	rows := make([][]int, len(m))
	for i, row := range m {
		rows[i] = make([]int, len(m[0]))
		copy(rows[i], row)
	}

	return rows
}

// Cols returns a list of matrix columns
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return [][]int{}
	}

	cols := make([][]int, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(m); j++ {
			cols[i][j] = m[j][i]
		}
	}

	return cols
}

// Set the value val at given row and column index
func (m Matrix) Set(r int, c int, val int) bool {
	if (r < 0 || len(m) <= r) || (c < 0 || len(m[0]) <= c) {
		return false
	}

	m[r][c] = val
	return true
}
