package matrix

import (
	"errors"
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
	var err error
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 1, []float64{1, 2, 3}, nil}
	matrix = NewVector(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
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

func TestErr(t *testing.T) {
	var matrix *Matrix
	matrix = New(1, 2, nil)
	matrix.err = errors.New("error")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
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

func TestAddRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{0, 0, 0, 0, 0, 0, 1, 2, 3}, nil}
	matrix = matrix.AddRow(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	matrix = matrix.AddRow(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRow(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 7, 7}, nil}
	matrix = matrix.AddRow(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7.6, 7.6, 7.6}, nil}
	matrix = matrix.AddRow(7.6)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRow("this is not allowed")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = New(2, 3, nil)
	vector = []float64{1, 2}
	matrix = matrix.AddRow(vector)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 2, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	matrix = matrix.AddRow(*matrix2)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestAddRowHEAD(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 0, 0, 0, 0, 0, 0}, nil}
	matrix = matrix.AddRowHEAD(vector)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err())
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}, nil}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{7, 7, 7, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	answer = &Matrix{3, 3, []float64{7.8, 7.8, 7.8, 1, 2, 3, 4, 5, 6}, nil}
	matrix = matrix.AddRowHEAD(7.8)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowHEADError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	matrix = New(2, 3, nil)
	vector = []float64{1, 2}
	matrix = matrix.AddRowHEAD(vector)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = New(2, 3, nil)
	matrix2 = New(2, 2, nil)
	matrix = matrix.AddRowHEAD(matrix2)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.AddRowHEAD("this type is not allowed")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func ExampleShow() {
	var matrix *Matrix
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix.Show()
	matrix = &Matrix{3, 2, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix.Show()
	// Output:
	// size: 2 x 3
	// 1 2 3
	// 4 5 6
	// size: 3 x 2
	// 1 2
	// 3 4
	// 5 6
}

func TestSize(t *testing.T) {
	var matrix *Matrix
	matrix = New(2, 3, nil)
	r, c := matrix.Size()
	if r != 2 || c != 3 {
		t.Errorf("want %v %v got %v %v", 2, 3, r, c)
	}
}

func TestRow(t *testing.T) {
	var matrix *Matrix
	matrix = New(7, 12, nil)
	r := matrix.Row()
	if r != 7 {
		t.Errorf("want %v got %v", 7, r)
	}
}

func TestColumn(t *testing.T) {
	var matrix *Matrix
	matrix = New(7, 12, nil)
	c := matrix.Column()
	if c != 12 {
		t.Errorf("want %v got %v", 12, c)
	}
}

func TestAt(t *testing.T) {
	var err error
	var matrix *Matrix
	matrix = New(2, 3, nil)
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
	matrix = New(2, 3, nil)
	answer = New(2, 3, nil)
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
	var matrix *Matrix
	var answer *Matrix
	matrix = createUniformMatrix(2, 3, 4)
	answer = createUniformMatrix(3, 5, 6)
	matrix.SetMatrix(answer)
	if matrix.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix.Err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	answer = &Matrix{-1, 2, []float64{}, nil}
	matrix.SetMatrix(answer)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
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

func TestSepRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	answer = &Matrix{3, 3, []float64{4, 5, 6, 7, 8, 9, 10, 11, 12}, nil}
	matrix2 = matrix.SepRow(2, 4)
	if matrix2.Err() != nil {
		t.Errorf("Should be error nil but got %v", matrix2.Err())
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2 = matrix.SepRow(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepRowError(t *testing.T) {
	var matrix *Matrix
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	matrix = matrix.SepRow(2, 1)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.SepRow(0, 4)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = matrix.SepRow(1, 7)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestSepColumn(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var err error
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	answer = &Matrix{3, 3, []float64{2, 3, 4, 8, 9, 10, 14, 15, 16}, nil}
	matrix2 = matrix.SepColumn(2, 4)
	if matrix2.Err() != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2 = matrix.SepColumn(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepColumnError(t *testing.T) {
	var matrix *Matrix
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, nil}
	matrix = matrix.SepColumn(2, 1)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = matrix.SepColumn(0, 4)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix = matrix.SepColumn(1, 7)
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
	}
}
