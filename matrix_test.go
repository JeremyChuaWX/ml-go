package mlgo

import (
	"errors"
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

func TestMultiply(t *testing.T) {
	test1 := &matrix{
		Rows: 3,
		Cols: 2,
		Data: [][]float64{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	}

	test2 := &matrix{
		Rows: 2,
		Cols: 3,
		Data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	test3 := &matrix{
		Rows: 3,
		Cols: 3,
		Data: [][]float64{
			{17, 22, 27},
			{22, 29, 36},
			{27, 36, 45},
		},
	}

	expected := &matrix{
		Rows: 3,
		Cols: 3,
		Data: [][]float64{
			{17, 22, 27},
			{22, 29, 36},
			{27, 36, 45},
		},
	}

	_, err := Multiply(test1, test3)

	if !errors.Is(err, shapeError) {
		t.Errorf("Multiply(%v, %v); want shape error", test1, test3)
	}

	result, _ := Multiply(test1, test2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"Multiply(%v, %v) = %v; want %v",
			test1, test2, result, expected,
		)
	}
}
