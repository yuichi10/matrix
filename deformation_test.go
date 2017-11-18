package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, nil}
	answer = &Matrix{4, 3, []float64{1, 5, 9, 2, 6, 10, 3, 7, 11, 4, 8, 12}, nil}
	matrix = matrix.Transpose()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestVector(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{6, 1, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = matrix.Vector()
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestReshape(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = NewVector([]float64{1, 2, 3, 4, 5, 6})
	answer = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.Reshape(2, 3)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestReshapeError(t *testing.T) {
	var matrix *Matrix
	matrix = NewVector([]float64{1, 2, 3, 4, 5, 6})
	matrix = matrix.Reshape(3, 3)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = NewVector([]float64{1, 2, 3, 4, 5, 6})
	matrix = matrix.Reshape(2, 2)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}
