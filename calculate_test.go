package matrix

import (
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
	matrix, _ = matrix.Add(5)
	answer = createUniformMatrix(5, 3, 6.3)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 1)
	matrix, _ = matrix.Add(3.4)
	answer = createUniformMatrix(5, 3, 4.4)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 2)
	matrix2 = createUniformMatrix(5, 3, 5)
	matrix, _ = matrix.Add(matrix2)
	answer = createUniformMatrix(5, 3, 7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 1)
	matrix2 = createUniformMatrix(5, 3, 4.7)
	matrix, _ = matrix.Add(*matrix2)
	answer = createUniformMatrix(5, 3, 5.7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	_, err = matrix.Add(matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
	_, err = matrix.Add(*matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}

	_, err = matrix.Add("not allowed type")
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
}

func TestSub(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(5, 3, 0)
	matrix, _ = matrix.Sub(4.3)
	answer = createUniformMatrix(5, 3, -4.3)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 5)
	matrix, _ = matrix.Sub(matrix2)
	answer = createUniformMatrix(5, 3, -5)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix2 = createUniformMatrix(5, 3, -5.3)
	matrix, _ = matrix.Sub(*matrix2)
	answer = createUniformMatrix(5, 3, 5.3)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = createUniformMatrix(5, 3, 0)
	matrix, _ = matrix.Sub(-3)
	answer = createUniformMatrix(5, 3, 3)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestSubError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createUniformMatrix(4, 3, 0)
	matrix2 = createUniformMatrix(5, 3, 1)
	_, err = matrix.Sub(matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
	_, err = matrix.Sub(*matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}

	_, err = matrix.Sub("not allowed type")
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
}

func TestMulti(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var matrix3 *Matrix
	var answer *Matrix
	var err error
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(4, 3, 0)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6, 7, 8}
	matrix2.matrix = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	matrix3, err = matrix.Multi(matrix2)
	answer = &Matrix{2, 3, []float64{150, 160, 170, 366, 392, 418}}
	if err != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}

	matrix = createUniformMatrix(2, 4, 3)
	matrix3, _ = matrix.Multi(3)
	answer = createUniformMatrix(2, 4, 9)
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}

	matrix = createUniformMatrix(2, 4, 3)
	matrix3, _ = matrix.Multi(3.5)
	answer = createUniformMatrix(2, 4, 10.5)
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}
}

func TestMultiError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(5, 3, 0)
	_, err = matrix.Multi(matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
	_, err = matrix.Multi(*matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix.Multi("string is not allowd")
	if err == nil {
		t.Errorf("opt2 is not allowed type therefore should get error but got success")
	}
}

func TestMultiEach(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var matrix3 *Matrix
	var answer *Matrix
	var err error
	matrix = &Matrix{3, 5, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}
	matrix2 = &Matrix{3, 5, []float64{1, 2, -3, 4, 5, 6, 7, -8, 9, 10, 11, 12, -13, 14, 15}}
	answer = &Matrix{3, 5, []float64{1, 4, -9, 16, 25, 36, 49, -64, 81, 100, 121, 144, -169, 196, 225}}
	matrix3, err = matrix.MultiEach(matrix2)
	if err != nil {
		t.Errorf("Matrix should be success but got error")
	}
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}

	matrix = createUniformMatrix(2, 4, 2)
	answer = &Matrix{2, 4, []float64{6, 6, 6, 6, 6, 6, 6, 6}}
	matrix3, _ = matrix.MultiEach(3)
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}

	matrix = createUniformMatrix(2, 4, 2)
	answer = &Matrix{2, 4, []float64{6.6, 6.6, 6.6, 6.6, 6.6, 6.6, 6.6, 6.6}}
	matrix3, _ = matrix.MultiEach(3.3)
	if !reflect.DeepEqual(answer, matrix3) {
		t.Errorf("want %#v got %#v", answer, matrix3)
	}
}

func TestMultiEachError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createUniformMatrix(2, 4, 0)
	matrix2 = createUniformMatrix(2, 3, 0)
	_, err = matrix.MultiEach(matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
	matrix2 = createUniformMatrix(3, 4, 0)
	_, err = matrix.MultiEach(*matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix.MultiEach("string is not allowd")
	if err == nil {
		t.Errorf("opt2 is not allowed type therefore should get error but got success")
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
