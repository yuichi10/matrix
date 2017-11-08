package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var err error
	matrix, err = NewMatrix(2, 3)
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}}
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	_, err = NewMatrix(-1, 2)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = NewMatrix(1, -2)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = NewMatrix(0, 2)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = NewMatrix(2, 0)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestNewVector(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var vector []float64
	var err error
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 1, []float64{1, 2, 3}}
	matrix, err = NewVector(vector)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix, err = NewVector(nil)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix, err = NewVector([]float64{})
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestCopy(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var err error
	answer = createUniformMatrix(2, 3, 5)
	matrix, err = Copy(answer)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	answer = &Matrix{2, 3, []float64{1, 2}}
	_, err = Copy(answer)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestZeroMatrix(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix, _ = NewMatrix(2, 3)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6}
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}}
	matrix.ZeroMatrix()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRow(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var vector []float64
	var err error
	matrix, _ = NewMatrix(2, 3)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{0, 0, 0, 0, 0, 0, 1, 2, 3}}
	err = matrix.AddRow(vector)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	matrix, _ = NewMatrix(2, 3)
	vector = []float64{1, 2}
	err = matrix.AddRow(vector)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestAddRowHEAD(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var vector []float64
	var err error
	matrix, _ = NewMatrix(2, 3)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 0, 0, 0, 0, 0, 0}}
	err = matrix.AddRowHEAD(vector)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	matrix, _ = NewMatrix(2, 3)
	vector = []float64{1, 2}
	err = matrix.AddRowHEAD(vector)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func ExampleShow() {
	var matrix *Matrix
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix.Show()
	matrix = &Matrix{3, 2, []float64{1, 2, 3, 4, 5, 6}}
	matrix.Show()
	// Output:
	// 1 2 3
	// 4 5 6
	// 1 2
	// 3 4
	// 5 6
}

func TestSize(t *testing.T) {
	var matrix *Matrix
	matrix, _ = NewMatrix(2, 3)
	r, c := matrix.Size()
	if r != 2 || c != 3 {
		t.Errorf("want %v %v got %v %v", 2, 3, r, c)
	}
}

func TestAt(t *testing.T) {
	var err error
	var matrix *Matrix
	matrix, _ = NewMatrix(2, 3)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6}
	count := 0
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			if val, _ := matrix.At(i, j); val != matrix.matrix[count] {
				t.Errorf("At(%v, %v) should be %v but got %v", i, j, count, val)
			}
			count++
		}
	}

	_, err = matrix.At(2, 4)
	if err == nil {
		t.Errorf("At(2, 4) is out of range thus should return error")
	}

	_, err = matrix.At(3, 3)
	if err == nil {
		t.Errorf("At(3, 3) is out of range thus should return error")
	}
}

func TestSet(t *testing.T) {
	var err error
	var matrix *Matrix
	var answer *Matrix
	matrix, _ = NewMatrix(2, 3)
	answer, _ = NewMatrix(2, 3)
	answer.matrix = []float64{1, 2, 3, 4, 5, 6}
	count := 0
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			count++
			matrix.Set(i, j, float64(count))
		}
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	err = matrix.Set(2, 4, 1)
	if err == nil {
		t.Errorf("At(2, 4) is out of range thus should return error")
	}

	err = matrix.Set(3, 3, 1)
	if err == nil {
		t.Errorf("At(3, 3) is out of range thus should return error")
	}
}

func TestSetMatrix(t *testing.T) {
	var err error
	var matrix *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(2, 3, 4)
	answer = createUniformMatrix(3, 5, 6)
	err = matrix.SetMatrix(answer)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	answer = &Matrix{-1, 2, []float64{}}
	err = matrix.SetMatrix(answer)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestTranspose(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}}
	answer = &Matrix{4, 3, []float64{1, 5, 9, 2, 6, 10, 3, 7, 11, 4, 8, 12}}
	matrix.Transpose()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}
