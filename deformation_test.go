package matrix

import (
	"math"
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

func TestInverse(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = New(4, 4, []float64{1, 2, 0, -1, -1, 1, 2, 0, 2, 0, 1, 1, 1, -2, -1, 1})
	answer = New(4, 4, []float64{
		1.9999999999999996, 2, -1, 3,
		-3.9999999999999987, -4.999999999999998, 3, -6.999999999999999,
		2.999999999999999, 3.999999999999999, -2, 4.999999999999999,
		-6.9999999999999964, -7.9999999999999964, 4.999999999999999, -10.999999999999996,
	})
	matrix = matrix.Inverse()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	matrix = matrix.Inverse()
	if !math.IsNaN(matrix.matrix[0]) {
		t.Errorf("want NaN got %v", matrix.matrix[0])
	}

	matrix = New(4, 4, []float64{3, 1, 1, 2, 5, 1, 3, 4, 2, 0, 0, 0, 1, 3, 2, 1})
	answer = New(4, 4, []float64{
		0, 0, 0.5, 0,
		0.6250000000000001, -0.3750000000000001, -0.1250000000000001, 0.25,
		-1.3749999999999996, 0.6249999999999999, 0.375, 0.24999999999999994,
		0.8749999999999994, -0.12499999999999986, -0.8749999999999999, -0.24999999999999994,
	})
	matrix = matrix.Inverse()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(4, 4, []float64{0, 1, 1, 2, 5, 1, 3, 4, 2, 0, 2, 0, 1, 3, 2, 1})
	answer = New(4, 4, []float64{
		-0.5789473684210524, 0.2631578947368422, -0.21052631578947384, 0.10526315789473674,
		-0.3157894736842105, 0.052631578947368474, -0.3421052631578948, 0.4210526315789473,
		0.5789473684210524, -0.2631578947368422, 0.7105263157894738, -0.10526315789473675,
		0.3684210526315788, 0.10526315789473681, -0.1842105263157894, -0.1578947368421052,
	})
	matrix = matrix.Inverse()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}
