package mlgo

import (
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	test := &matrix{
		Rows: 2,
		Cols: 2,
		Data: [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	result := &matrix{
		Rows: 2,
		Cols: 2,
		Data: [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	expected := &matrix{
		Rows: 2,
		Cols: 2,
		Data: [][]float64{
			{2, 4},
			{6, 8},
		},
	}

	Scale(result, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Scale(%v, 2) = %v; expected %v", test, result, expected)
	}
}

func TestTranspose(t *testing.T) {
	test := &matrix{
		Rows: 2,
		Cols: 3,
		Data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	expected := &matrix{
		Rows: 3,
		Cols: 2,
		Data: [][]float64{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	}

	result := Transpose(test)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Transpose(%v) = %v; want %v", test, result, expected)
	}
}
