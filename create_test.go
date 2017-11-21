package matrix

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var err error
	matrix = New(2, 3, nil)
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}, nil}
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(2, 3, []float64{1, 2, 3, 4, 5, 6})
	answer = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(2, 3, []float64{})
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}, nil}
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(2, 3, 7)
	answer = &Matrix{2, 3, []float64{7, 7, 7, 7, 7, 7}, nil}
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = New(2, 3, 7.7)
	answer = &Matrix{2, 3, []float64{7.7, 7.7, 7.7, 7.7, 7.7, 7.7}, nil}
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestNewError(t *testing.T) {
	var matrix *Matrix
	matrix = New(-1, 2, nil)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(1, -2, nil)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(0, 2, nil)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(2, 0, nil)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = New(2, 3, []float64{1, 2, 3, 4, 5, 6, 7})
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(2, 3, []float64{1, 2, 3, 4, 5})
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(2, 3, "not allowed argument")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestNewVector(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var vector []float64
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 1, []float64{1, 2, 3}, nil}
	matrix = NewVector(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = NewVector(nil)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = NewVector([]float64{})
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestNewRandom(t *testing.T) {
	var matrix *Matrix
	matrix = NewRandom(3, 4, 5)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 4; j++ {
			val, _ := matrix.At(i, j)
			if val > 1 || val < 0 {
				t.Errorf("should be 0~1")
			}
		}
	}
}

func TestNewRandomError(t *testing.T) {
	var matrix *Matrix
	matrix = NewRandom(0, 2, 2)
	if matrix.Err() == nil {
		t.Errorf("You should get error but got nil")
	}
}

func TestNewHotVector(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = NewHotVector(5, 3)
	answer = NewVector([]float64{0, 0, 1, 0, 0})
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestNewHotVectorError(t *testing.T) {
	var matrix *Matrix
	matrix = NewHotVector(3, 0)
	if matrix.Err() == nil {
		t.Errorf("You should get error but got nil")
	}

	matrix = NewHotVector(3, 4)
	if matrix.Err() == nil {
		t.Errorf("You should get error but got nil")
	}
}

func TestNewEye(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = NewEye(3)
	answer = New(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestNewEyeError(t *testing.T) {
	var matrix *Matrix
	matrix = NewEye(0)
	if matrix.Err() == nil {
		t.Errorf("You should get error but got nil")
	}

	matrix = NewEye(-1)
	if matrix.Err() == nil {
		t.Errorf("You should get error but got nil")
	}
}

func TestCopy(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	answer = createUniformMatrix(2, 3, 5)
	matrix = Copy(answer)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}
