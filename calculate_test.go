package matrix

import "testing"

func createMatrix(row, column int, val float64) *Matrix {
	matrix := NewMatrix(row, column)
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
