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

func TestSigmoid(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, -1, -2, -3}, nil}
	answer = &Matrix{3, 3, []float64{731059, 880797, 952574, 982014, 993307, 997527, 268941, 119203, 47426}, nil}
	matrix2 = matrix.Sigmoid()
	matrix2 = matrix2.MultiEach(1000000)
	for i := 1; i <= matrix2.row; i++ {
		for j := 1; j <= matrix2.column; j++ {
			val, _ := matrix2.At(i, j)
			ans, _ := answer.At(i, j)
			if int(math.Floor(val+.5)) != int(ans) {
				t.Errorf("want %#v got %#v", int(ans), math.Floor(val+.5))
			}
		}
	}
}

func TestSigmoidGradient(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, -1, -2, -3}, nil}
	answer = &Matrix{3, 3, []float64{1966119, 1049936, 451767, 176627, 66481, 24665, 1966119, 1049936, 451767}, nil}
	matrix = matrix.SigmoidGradient().MultiEach(10000000)
	for i := 1; i <= matrix.row; i++ {
		for j := 1; j <= matrix.column; j++ {
			val, _ := matrix.At(i, j)
			ans, _ := answer.At(i, j)
			if int(math.Floor(val+.5)) != int(ans) {
				t.Errorf("want %#v got %#v", int(ans), math.Floor(val+.5))
			}
		}
	}
}
