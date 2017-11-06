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
