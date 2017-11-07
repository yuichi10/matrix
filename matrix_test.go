package matrix

import "testing"

func TestAt(t *testing.T) {
	var err error
	var matrix *Matrix
	matrix, _ = NewMatrix(2, 3)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6}
	count := 0
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			count++
			if val, _ := matrix.At(i, j); val != float64(count) {
				t.Errorf("At(%v, %v) should be %v but got %v", i, j, count, val)
			}
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
	matrix, _ = NewMatrix(2, 3)
	count := 0
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			count++
			matrix.Set(i, j, float64(count))
			if val, _ := matrix.At(i, j); val != float64(count) {
				t.Errorf("At(%v, %v) should be %v but got %v", i, j, count, val)
			}
		}
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
