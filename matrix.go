package mlgo

import "math/rand"

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

// univariate

func Scale(m *matrix, x float64) {
	for r := range m.Rows {
		for c := range m.Cols {
			m.Data[r][c] = x * m.Data[r][c]
		}
	}
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

// bivariate
