package matrix

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func createUniformMatrix(row, column int, val float64) *Matrix {
	matrix, _ := New(row, column, nil)
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			matrix.Set(i, j, val)
		}
	}
	return matrix
}

func TestAdd(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(5, 3, 1.3)
	matrix = matrix.Add(5)
	answer = createUniformMatrix(5, 3, 6.3)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 1)
	matrix = matrix.Add(3.4)
	answer = createUniformMatrix(5, 3, 4.4)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 2)
	matrix2 = createUniformMatrix(5, 3, 5)
	matrix = matrix.Add(matrix2)
	answer = createUniformMatrix(5, 3, 7)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 1)
	matrix2 = createUniformMatrix(5, 3, 4.7)
	matrix = matrix.Add(*matrix2)
	answer = createUniformMatrix(5, 3, 5.7)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix = matrix.Add(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix = matrix.Add(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix = matrix.Add("not allowed type")
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Add(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Add(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix.calcErr = errors.New("error")
	matrix = matrix.Add(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}
}

func TestSub(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(5, 3, 0)
	matrix = matrix.Sub(4.3)
	answer = createUniformMatrix(5, 3, -4.3)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 5)
	matrix = matrix.Sub(matrix2)
	answer = createUniformMatrix(5, 3, -5)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix2 = createUniformMatrix(5, 3, -5.3)
	matrix = matrix.Sub(*matrix2)
	answer = createUniformMatrix(5, 3, 5.3)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix = matrix.Sub(-3)
	answer = createUniformMatrix(5, 3, 3)
	if matrix.CalcErr() != nil {
		t.Errorf("Should be error nil but got %v", matrix.CalcErr())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestSubError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix = matrix.Sub(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix = matrix.Sub(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix = matrix.Sub("not allowed type")
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Sub(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Sub(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	matrix.calcErr = errors.New("error")
	matrix = matrix.Sub(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("Sub should get error got nil")
	}
}

func TestMulti(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(4, 3, 0)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6, 7, 8}
	matrix2.matrix = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	matrix = matrix.Multi(matrix2)
	answer = &Matrix{2, 3, []float64{150, 160, 170, 366, 392, 418}, nil}
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(2, 4, 3)
	matrix = matrix.Multi(3)
	answer = createUniformMatrix(2, 4, 9)
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(2, 4, 3)
	matrix = matrix.Multi(3.5)
	answer = createUniformMatrix(2, 4, 10.5)
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestMultiError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(5, 3, 0)
	matrix = matrix.Multi(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
	matrix = createUniformMatrix(2, 4, 0)
	matrix = matrix.Multi(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
	matrix = createUniformMatrix(2, 4, 0)
	matrix = matrix.Multi("string is not allowd")
	if matrix.CalcErr() == nil {
		t.Errorf("opt2 is not allowed type therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(5, 3, 0)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Multi(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(5, 3, 0)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Multi(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(5, 3, 0)
	matrix.calcErr = errors.New("error")
	matrix = matrix.Multi(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
}

func TestMultiEach(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 5, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	matrix2 = &Matrix{3, 5, []float64{1, 2, -3, 4, 5, 6, 7, -8, 9, 10, 11, 12, -13, 14, 15}, nil}
	answer = &Matrix{3, 5, []float64{1, 4, -9, 16, 25, 36, 49, -64, 81, 100, 121, 144, -169, 196, 225}, nil}
	matrix = matrix.MultiEach(matrix2)
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(2, 4, 2)
	answer = &Matrix{2, 4, []float64{6, 6, 6, 6, 6, 6, 6, 6}, nil}
	matrix = matrix.MultiEach(3)
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(2, 4, 2)
	answer = &Matrix{2, 4, []float64{6.6, 6.6, 6.6, 6.6, 6.6, 6.6, 6.6, 6.6}, nil}
	matrix = matrix.MultiEach(3.3)
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestMultiEachError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(2, 3, 0)
	matrix = matrix.MultiEach(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(3, 4, 0)
	matrix = matrix.MultiEach(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix = matrix.MultiEach("string is not allowd")
	if matrix.CalcErr() == nil {
		t.Errorf("opt2 is not allowed type therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(3, 4, 0)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.MultiEach(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(3, 4, 0)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.MultiEach(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(3, 4, 0)
	matrix.calcErr = errors.New("error")
	matrix = matrix.MultiEach(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
}

func TestDiv(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix, _ = New(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	matrix2, _ = New(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	matrix = matrix.Div(matrix2)
	answer = &Matrix{2, 4, []float64{1, 1, 1, 1, 1, 1, 1, 1}, nil}
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix, _ = New(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	matrix2, _ = New(2, 4, []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8})
	matrix = matrix.Div(*matrix2)
	answer = &Matrix{2, 4, []float64{10, 10, 10, 10, 10, 10, 10, 10}, nil}
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix, _ = New(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	matrix = matrix.Div(2)
	answer = &Matrix{2, 4, []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4}, nil}
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix, _ = New(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	matrix = matrix.Div(0.1)
	answer = &Matrix{2, 4, []float64{10, 20, 30, 40, 50, 60, 70, 80}, nil}
	if matrix.CalcErr() != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestDivError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix, _ = New(2, 4, 2)
	matrix2, _ = New(2, 3, 3)
	matrix = matrix.Div(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix, _ = New(2, 4, 2)
	matrix2, _ = New(2, 3, 3)
	matrix = matrix.Div(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix, _ = New(2, 4, 2)
	matrix = matrix.Div("string is not allowd")
	if matrix.CalcErr() == nil {
		t.Errorf("opt2 is not allowed type therefore should get error but got success")
	}

	matrix, _ = New(4, 3, 0)
	matrix2, _ = New(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Div(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt2 is error thus should get error got nil")
	}

	matrix, _ = New(4, 3, 0)
	matrix2, _ = New(5, 3, 1)
	matrix2.calcErr = errors.New("error")
	matrix = matrix.Div(*matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt2 is error thus should get error got nil")
	}

	matrix, _ = New(4, 3, 0)
	matrix2, _ = New(5, 3, 1)
	matrix.calcErr = errors.New("error")
	matrix = matrix.Div(matrix2)
	if matrix.CalcErr() == nil {
		t.Errorf("opt1 is error thus should get error got nil")
	}
}

func setup() {}

func teardown() {}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}
