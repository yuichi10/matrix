package matrix

import (
	"errors"
	"reflect"
	"testing"
)

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

func TestSize(t *testing.T) {
	var matrix *Matrix
	matrix = New(2, 3, nil)
	r, c := matrix.Size()
	if r != 2 || c != 3 {
		t.Errorf("want %v %v got %v %v", 2, 3, r, c)
	}
}

func ExampleShow() {
	var matrix *Matrix
	matrix = &Matrix{2, 3, []float64{1, 2, 3, 4, 5, 6}, nil}
	matrix.Show()
	matrix = &Matrix{3, 2, []float64{1.1, 2.1234, 3.12345, 4.123456, 5.123455, 0.6}, nil}
	matrix.err = newError("Something happen", "Show", matrix, nil)
	matrix.Show()
	// Output:
	// size: 2 x 3
	// 1.00000 2.00000 3.00000
	// 4.00000 5.00000 6.00000
	//
	// size: 3 x 2
	// 1.10000 2.12340
	// 3.12345 4.12346
	// 5.12345 0.60000
	// Size: opt1: (3, 2), opt2: No matrix
	// [function: Show] Something happen
}

func TestErr(t *testing.T) {
	var matrix *Matrix
	matrix = New(1, 2, nil)
	matrix.err = errors.New("error")
	if matrix.Err() == nil {
		t.Errorf("Should get error but got nil")
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
		t.Errorf("Should be error nil but got %v", matrix.Err())
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

func TestInt(t *testing.T) {
	var matrix *Matrix
	var answer *Matrix
	matrix = New(3, 3, []float64{0.4, 0.5, 0.6, 1.4, 1.5, 1.6, 2.4, 2.5, 2.9})
	answer = New(3, 3, []float64{0, 0, 0, 1, 1, 1, 2, 2, 2})
	matrix = matrix.Int()
	if !reflect.DeepEqual(answer, matrix) {
		t.Errorf("want %#v got %#v", answer, matrix)
	}
}

func TestTextAsOneLine(t *testing.T) {
	var matrix *Matrix
	var text string
	matrix = New(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	text = matrix.TextAsOneLine(",")
	if text != "1.000000,2.000000,3.000000,4.000000,5.000000,6.000000,7.000000,8.000000,9.000000,10.000000,11.000000,12.000000" {
		t.Errorf("want %q got %q", "1.000000,2.000000,3.000000,4.000000,5.000000,6.000000,7.000000,8.000000,9.000000,10.000000,11.000000,12.000000", text)
	}

	matrix = New(4, 1, []float64{1, 2, 3, 4})
	text = matrix.TextAsOneLine(",")
	if text != "1.000000,2.000000,3.000000,4.000000" {
		t.Errorf("want %q got %q", "1.000000,2.000000,3.000000,4.000000", text)
	}

	matrix = New(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	text = matrix.TextAsOneLine(" ")
	if text != "1.000000 2.000000 3.000000 4.000000 5.000000 6.000000 7.000000 8.000000 9.000000 10.000000 11.000000 12.000000" {
		t.Errorf("want %q got %q", "1.000000 2.000000 3.000000 4.000000 5.000000 6.000000 7.000000 8.000000 9.000000 10.000000 11.000000 12.000000", text)
	}
}
