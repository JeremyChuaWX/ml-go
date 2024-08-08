package mlgo

import (
	"errors"
	"fmt"
	"math/rand"
)

var shapeError = errors.New("matrix shapes are incompatible")

var valueMismatchError = errors.New("value mismatch")

type matrix struct {
	Rows int
	Cols int
	Data [][]float64 // row by column
}

func ZeroMatrix(rows int, cols int) *matrix {
	data := make([][]float64, rows)
	for i := range rows {
		data[i] = make([]float64, cols)
	}
	return &matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func RandomMatrix(rows int, cols int) *matrix {
	data := make([][]float64, rows)
	for r := range rows {
		data[r] = make([]float64, cols)
		for c := range cols {
			data[r][c] = 1 - rand.Float64() // (0.0, 1.0]
		}
	}
	return &matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func (m *matrix) Scale(x float64) {
	for r := range m.Rows {
		for c := range m.Cols {
			m.Data[r][c] = x * m.Data[r][c]
		}
	}
}

func (m *matrix) Add(n *matrix) error {
	if m.Cols != n.Cols || m.Rows != n.Rows {
		return fmt.Errorf("%w: m=%v, n=%v", shapeError, m, n)
	}
	for r := range m.Rows {
		for c := range m.Cols {
			m.Data[r][c] += n.Data[r][c]
		}
	}
	return nil
}

func (m *matrix) Subtract(n *matrix) error {
	if m.Cols != n.Cols || m.Rows != n.Rows {
		return fmt.Errorf("%w: m=%v, n=%v", shapeError, m, n)
	}
	for r := range m.Rows {
		for c := range m.Cols {
			m.Data[r][c] -= n.Data[r][c]
		}
	}
	return nil
}

// Hadamard product (element-wise product)
func (m *matrix) Multiply(n *matrix) error {
	if m.Cols != n.Cols || m.Rows != n.Rows {
		return fmt.Errorf("%w: m=%v, n=%v", shapeError, m, n)
	}
	for r := range m.Rows {
		for c := range m.Cols {
			m.Data[r][c] *= n.Data[r][c]
		}
	}
	return nil
}

func Copy(m *matrix) *matrix {
	cpy := ZeroMatrix(m.Rows, m.Cols)
	for r := range m.Rows {
		for c := range m.Cols {
			cpy.Data[r][c] = m.Data[r][c]
		}
	}
	return cpy
}

func Transpose(m *matrix) *matrix {
	mp := ZeroMatrix(m.Cols, m.Rows)
	for r := range m.Rows {
		for c := range m.Cols {
			mp.Data[c][r] = m.Data[r][c]
		}
	}
	return mp
}

func RowSum(m *matrix) *matrix {
	res := ZeroMatrix(m.Rows, 1)
	for r := range m.Rows {
		for c := range m.Cols {
			res.Data[r][0] += m.Data[r][c]
		}
	}
	return res
}

func ColSum(m *matrix) *matrix {
	res := ZeroMatrix(1, m.Cols)
	for r := range m.Rows {
		for c := range m.Cols {
			res.Data[0][c] += m.Data[r][c]
		}
	}
	return res
}

func Equal(m *matrix, n *matrix) (bool, error) {
	if m.Cols != n.Cols || m.Rows != n.Rows {
		return false, fmt.Errorf("%w: m=%v, n=%v", shapeError, m, n)
	}
	for r := range m.Rows {
		for c := range m.Cols {
			if m.Data[r][c] != n.Data[r][c] {
				return false, fmt.Errorf(
					"%w (%d, %d): m=%v, n=%v",
					valueMismatchError, r, c, m.Data[r][c], n.Data[r][c],
				)
			}
		}
	}
	return true, nil
}

func Dot(m *matrix, n *matrix) (*matrix, error) {
	if m.Cols != n.Rows {
		return nil, fmt.Errorf("%w: m=%v, n=%v", shapeError, m, n)
	}
	result := ZeroMatrix(m.Rows, n.Cols)
	for r := range result.Rows {
		for c := range result.Cols {
			for i := 0; i < m.Cols; i++ {
				result.Data[r][c] += m.Data[r][i] * n.Data[i][c]
			}
		}
	}
	return result, nil
}
