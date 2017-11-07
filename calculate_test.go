package matrix

import (
	"os"
	"testing"
)

func createMatrix(row, column int, val float64) *Matrix {
	matrix, _ := NewMatrix(row, column)
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			matrix.Set(i, j, val)
		}
	}
	return matrix
}

func TestAdd(t *testing.T) {
	matrix := createMatrix(5, 3, 0)
	matrix.Add(5)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != 5 {
				t.Errorf("Add was incorrect, got: %v, want: %v.", result, 5)
			}
		}
	}

	matrix = createMatrix(5, 3, 0)
	matrix2 := createMatrix(5, 3, 5)
	matrix.Add(matrix2)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != 5 {
				t.Errorf("Add was incorrect, got: %v, want: %v.", result, 5)
			}
		}
	}

	matrix = createMatrix(5, 3, 0)
	matrix2 = createMatrix(5, 3, 5)
	matrix.Add(*matrix2)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != 5 {
				t.Errorf("Add was incorrect, got: %v, want: %v.", result, 5)
			}
		}
	}
}

func TestAddError(t *testing.T) {
	matrix := createMatrix(5, 3, 0)
	matrix2 := createMatrix(5, 2, 5)
	err := matrix.Add(matrix2)
	if err == nil {
		t.Errorf("Add should get error, got: %v, want: error", err)
	}

	matrix2 = createMatrix(4, 3, 0)
	err = matrix.Add(matrix2)
	if err == nil {
		t.Errorf("Add should get error, got: %v, want: error", err)
	}
}

func TestSub(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createMatrix(5, 3, 0)
	matrix.Sub(4.3)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != -4.3 {
				t.Errorf("Sub was incorrect, got: %v, want: %v.", result, -4.3)
			}
		}
	}

	matrix = createMatrix(5, 3, 0)
	matrix2 = createMatrix(5, 3, 5)
	matrix.Sub(matrix2)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != -5 {
				t.Errorf("Sub was incorrect, got: %v, want: %v.", result, -5)
			}
		}
	}

	matrix = createMatrix(5, 3, 0)
	matrix2 = createMatrix(5, 3, -5.3)
	matrix.Sub(*matrix2)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != 5.3 {
				t.Errorf("Sub was incorrect, got: %v, want: %v.", result, 5.3)
			}
		}
	}

	matrix = createMatrix(5, 3, 0)
	matrix.Sub(-3)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			result, _ := matrix.At(i, j)
			if result != 3 {
				t.Errorf("Sub was incorrect, got: %v, want: %v.", result, 3)
			}
		}
	}
}

func TestSubError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	matrix = createMatrix(4, 3, 0)
	matrix.rows = -3
	err := matrix.Sub(4.3)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}

	matrix = createMatrix(4, 3, 0)
	matrix2 = createMatrix(5, 3, 1)
	err = matrix.Sub(matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
	err = matrix.Sub(*matrix2)
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}

	err = matrix.Sub("not allowed type")
	if err == nil {
		t.Errorf("Sub should get error got nil")
	}
}

func TestMulti(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createMatrix(2, 4, 0)
	matrix2 = createMatrix(4, 3, 0)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6, 7, 8}
	matrix2.matrix = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	matrix.Multi(matrix2)
	if matrix.rows != 2 {
		t.Errorf("Matrix row length should be 3 now %v", matrix.rows)
	}
	if matrix.columns != 3 {
		t.Errorf("Matrix column length should be 4 now %v", matrix.columns)
	}
	if val, _ := matrix.At(1, 1); val != 150 {
		t.Errorf("Matrix column at (1,1) should be 150 now %v", val)
	}
	if val, _ := matrix.At(1, 2); val != 160 {
		t.Errorf("Matrix column at (1,2) should be 160 now %v", val)
	}
	if val, _ := matrix.At(1, 3); val != 170 {
		t.Errorf("Matrix column at (1,3) should be 170 now %v", val)
	}
	if val, _ := matrix.At(2, 1); val != 366 {
		t.Errorf("Matrix column at (2,1) should be 366 now %v", val)
	}
	if val, _ := matrix.At(2, 2); val != 392 {
		t.Errorf("Matrix column at (2,2) should be 392 now %v", val)
	}
	if val, _ := matrix.At(2, 3); val != 418 {
		t.Errorf("Matrix column at (2,3) should be 418 now %v", val)
	}

	matrix = createMatrix(2, 4, 0)
	matrix2 = createMatrix(4, 3, 0)
	err = matrix.Multi(*matrix2)
	if err != nil {
		t.Errorf("Matrix should be success but got error")
	}

	matrix = createMatrix(2, 4, 3)
	matrix.Multi(3)
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 4; j++ {
			result, _ := matrix.At(i, j)
			if result != 9 {
				t.Errorf("All matrix value should be 9 but got %v", result)
			}
		}
	}

	matrix = createMatrix(2, 4, 3)
	matrix.Multi(3.5)
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 4; j++ {
			result, _ := matrix.At(i, j)
			if result != 10.5 {
				t.Errorf("All matrix value should be 10.5 but got %v", result)
			}
		}
	}
}

func TestMultiError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var err error
	matrix = createMatrix(2, 4, 0)
	matrix2 = createMatrix(5, 3, 0)
	err = matrix.Multi(matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}
	err = matrix.Multi(*matrix2)
	if err == nil {
		t.Errorf("opt1's colmn and opt2's row is different therefore should get error but got success")
	}

	matrix.Multi("string is not allowd")
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
