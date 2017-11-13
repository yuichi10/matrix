package matrix

import (
	"math"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	var err error
	matrix, err = New(2, 3, nil)
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}}
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix, err = New(2, 3, []float64{1, 2, 3, 4, 5, 6})
	answer = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestNewError(t *testing.T) {
	var err error
	_, err = New(-1, 2, nil)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = New(1, -2, nil)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = New(0, 2, nil)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = New(2, 0, nil)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
	_, err = New(2, 3, []float64{1, 2, 3, 4, 5, 6, 7})
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = New(2, 3, []float64{1, 2, 3, 4, 5})
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
	answer = createUniformMatrix(2, 3, 5)
	matrix = Copy(answer)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestZeroMatrix(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix, _ = New(2, 3, nil)
	matrix.matrix = []float64{1, 2, 3, 4, 5, 6}
	answer = &Matrix{2, 3, []float64{0, 0, 0, 0, 0, 0}}
	matrix.ZeroMatrix()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	var err error
	matrix, _ = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{0, 0, 0, 0, 0, 0, 1, 2, 3}}
	matrix, err = matrix.AddRow(vector)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
	answer = &Matrix{5, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}
	matrix, _ = matrix.AddRow(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix, _ = matrix.AddRow(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 7, 7}}
	matrix, _ = matrix.AddRow(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, 7.6, 7.6, 7.6}}
	matrix, _ = matrix.AddRow(7.6)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	var err error
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	_, err = matrix.AddRow("this is not allowed")
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix, _ = New(2, 3, nil)
	vector = []float64{1, 2}
	matrix, err = matrix.AddRow(vector)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix2 = &Matrix{3, 2, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
	_, err = matrix.AddRow(*matrix2)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestAddRowHEAD(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var vector []float64
	var err error
	matrix, _ = New(2, 3, nil)
	vector = []float64{1, 2, 3}
	answer = &Matrix{3, 3, []float64{1, 2, 3, 0, 0, 0, 0, 0, 0}}
	matrix, err = matrix.AddRowHEAD(vector)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}}
	matrix, _ = matrix.AddRowHEAD(matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	matrix2 = &Matrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
	answer = &Matrix{5, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6}}
	matrix, _ = matrix.AddRowHEAD(*matrix2)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	answer = &Matrix{3, 3, []float64{7, 7, 7, 1, 2, 3, 4, 5, 6}}
	matrix, _ = matrix.AddRowHEAD(7)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}

	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	answer = &Matrix{3, 3, []float64{7.8, 7.8, 7.8, 1, 2, 3, 4, 5, 6}}
	matrix, _ = matrix.AddRowHEAD(7.8)
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestAddRowHEADError(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var vector []float64
	var err error
	matrix, _ = New(2, 3, nil)
	vector = []float64{1, 2}
	_, err = matrix.AddRowHEAD(vector)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
	matrix, _ = New(2, 3, nil)
	matrix2, _ = New(2, 2, nil)
	_, err = matrix.AddRowHEAD(matrix2)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = matrix.AddRowHEAD("this type is not allowed")
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
	matrix, _ = New(2, 3, nil)
	r, c := matrix.Size()
	if r != 2 || c != 3 {
		t.Errorf("want %v %v got %v %v", 2, 3, r, c)
	}
}

func TestRow(t *testing.T) {
	var matrix *Matrix
	matrix, _ = New(7, 12, nil)
	r := matrix.Row()
	if r != 7 {
		t.Errorf("want %v got %v", 7, r)
	}
}

func TestColumn(t *testing.T) {
	var matrix *Matrix
	matrix, _ = New(7, 12, nil)
	c := matrix.Column()
	if c != 12 {
		t.Errorf("want %v got %v", 12, c)
	}
}

func TestAt(t *testing.T) {
	var err error
	var matrix *Matrix
	matrix, _ = New(2, 3, nil)
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
	matrix, _ = New(2, 3, nil)
	answer, _ = New(2, 3, nil)
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
	matrix = matrix.Transpose()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestSepRow(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var err error
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	answer = &Matrix{3, 3, []float64{4, 5, 6, 7, 8, 9, 10, 11, 12}}
	matrix2, err = matrix.SepRow(2, 4)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2, err = matrix.SepRow(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepRowError(t *testing.T) {
	var matrix *Matrix
	var err error
	matrix = &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	_, err = matrix.SepRow(2, 1)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = matrix.SepRow(0, 4)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
	_, err = matrix.SepRow(1, 7)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestSepColumn(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	var err error
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	answer = &Matrix{3, 3, []float64{2, 3, 4, 8, 9, 10, 14, 15, 16}}
	matrix2, err = matrix.SepColumn(2, 4)
	if err != nil {
		t.Errorf("Should be error nil but got %v", err)
	}
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}

	answer = matrix
	matrix2, err = matrix.SepColumn(1, 6)
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSepColumnError(t *testing.T) {
	var matrix *Matrix
	var err error
	matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
	_, err = matrix.SepColumn(2, 1)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}

	_, err = matrix.SepColumn(0, 4)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
	_, err = matrix.SepColumn(1, 7)
	if err == nil {
		t.Errorf("Should get error but got nil")
	}
}

func TestVector(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}}
	answer = &Matrix{6, 1, []float64{1, 2, 3, 4, 5, 6}}
	matrix2 = matrix.Vector()
	if !reflect.DeepEqual(answer, matrix2) {
		t.Errorf("want %#v got %#v", answer, matrix2)
	}
}

func TestSigmoid(t *testing.T) {
	var matrix *Matrix
	var matrix2 *Matrix
	var answer *Matrix
	matrix = &Matrix{3, 3, []float64{1, 2, 3, 4, 5, 6, -1, -2, -3}}
	answer = &Matrix{3, 3, []float64{731059, 880797, 952574, 982014, 993307, 997527, 268941, 119203, 47426}}
	matrix2 = matrix.Sigmoid()
	matrix2, _ = matrix2.MultiEach(1000000)
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
