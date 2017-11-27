package matrix

import "testing"
import "fmt"

func TestPermutation(t *testing.T) {
	f := func(each []int, preRes *PermResult, arg interface{}) *PermResult {
		if preRes.value == nil {
			preRes.value = fmt.Sprintf("%s: %v", arg, each)
			return preRes
		}
		preRes.value = fmt.Sprintf("%s\n%s: %v", preRes.value, arg, each)
		return preRes
	}
	perm := PermutationProcess(3, f, "Each")
	answer := "Each: [1 2 3]\nEach: [1 3 2]\nEach: [3 1 2]\nEach: [2 1 3]\nEach: [2 3 1]\nEach: [3 2 1]"
	if perm.result.value.(string) != answer {
		t.Errorf("want %v but got %v", answer, perm.result.value)
	}
}

func TestSgn(t *testing.T) {
	result := 0
	result = Sgn([]int{1, 2, 3, 4})
	if result != 1 {
		t.Errorf("want %#v got %#v", 1, result)
	}

	result = Sgn([]int{2, 1, 3, 4})
	if result != -1 {
		t.Errorf("want %#v got %#v", -1, result)
	}

	result = Sgn([]int{2, 1, 4, 3})
	if result != 1 {
		t.Errorf("want %#v got %#v", 1, result)
	}

	result = Sgn([]int{4, 1, 2, 3})
	if result != -1 {
		t.Errorf("want %#v got %#v", -1, result)
	}

	result = Sgn(nil)
	if result != 0 {
		t.Errorf("want %#v got %#v", 0, result)
	}
}

func TestDeterminant(t *testing.T) {
	var matrix *Matrix
	var result float64
	var answer float64
	var err error
	matrix = New(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	result, err = matrix.Determinant()
	answer = 0
	if err != nil {
		t.Errorf("The error should be nil but got %v", err)
	}
	if result != answer {
		t.Errorf("want %v but got %v", answer, result)
	}

	matrix = New(4, 4, []float64{3, 1, 1, 2, 5, 1, 3, 4, 2, 0, 1, 0, 1, 3, 2, 1})
	result, err = matrix.Determinant()
	answer = -22
	if err != nil {
		t.Errorf("The error should be nil but got %v", err)
	}
	if result != answer {
		t.Errorf("want %v but got %v", answer, result)
	}
}

func TestDeterminantError(t *testing.T) {
	var matrix *Matrix
	var err error
	matrix = New(3, 2, []float64{1, 2, 3, 4, 5, 6})
	_, err = matrix.Determinant()
	if err == nil {
		t.Errorf("want error but got nil")
	}
}

func benchDeterminant(b *testing.B, count, size int) {
	matrix := NewRandom(size, size, 3)
	b.ReportAllocs()
	for i := 0; i < count; i++ {
		matrix.Determinant()
	}
}

func BenchmarkDeterminant5_5(b *testing.B)   { benchDeterminant(b, 5, 5) }
func BenchmarkDeterminant50_5(b *testing.B)  { benchDeterminant(b, 50, 5) }
func BenchmarkDeterminant5_10(b *testing.B)  { benchDeterminant(b, 5, 10) }
func BenchmarkDeterminant50_10(b *testing.B) { benchDeterminant(b, 50, 10) }
