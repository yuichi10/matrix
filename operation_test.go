package matrix

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = New(2, 3, nil)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6}
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}, nil}
	matrix = matrix.ZeroMatrix()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{0, 0, 0, 0, 0, 0, 1, 2, 3}, nil}
	matrix = matrix.AddRow(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	matrix = matrix.AddRow(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRow(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 7, 7}, nil}
	matrix = matrix.AddRow(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7.6, 7.6, 7.6}, nil}
	matrix = matrix.AddRow(7.6)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRow("this is not allowed")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(2, 3, nil)
	vector = []float64{1, 2}
	matrix = matrix.AddRow(vector)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 2, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	matrix = matrix.AddRow(*matrix2)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestAddRowHEAD(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 0, 0, 0, 0, 0, 0}, nil}
	matrix = matrix.AddRowHEAD(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{7, 7, 7, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{7.8, 7.8, 7.8, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(7.8)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowHEADError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2}
	matrix = matrix.AddRowHEAD(vector)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = New(2, 3, nil)
	matrix2 = New(2, 2, nil)
	matrix = matrix.AddRowHEAD(matrix2)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.AddRowHEAD("this type is not allowed")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestSepRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	answer = &Matrix{3, 3, []float64{4, 5, 6, 7, 8, 9, 10, 11, 12}, nil}
	matrix2 = matrix.SepRow(2, 4)
	if matrix2.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix2.Err())
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2 = matrix.SepRow(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepRowError(t *testing.T) {
	var matrix *Matrix
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	matrix = matrix.SepRow(2, 1)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.SepRow(0, 4)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = matrix.SepRow(1, 7)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestSepColumn(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var err error
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	answer = &Matrix{3, 3, []float64{2, 3, 4, 8, 9, 10, 14, 15, 16}, nil}
	matrix2 = matrix.SepColumn(2, 4)
	if matrix2.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2 = matrix.SepColumn(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepColumnError(t *testing.T) {
	var matrix *Matrix
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	matrix = matrix.SepColumn(2, 1)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.SepColumn(0, 4)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = matrix.SepColumn(1, 7)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}
