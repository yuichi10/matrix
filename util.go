package matrix

import (
	"errors"
	"math"
)

var a []int
var n int
var count int

// PermEachResult have the answer of return value of each perm calc result
type PermResult struct {
	value interface{}
	err   error
}

// Perm will have every infomation to do permutation
type Perm struct {
	length   int
	f        func([]int, *PermResult, interface{}) *PermResult
	argument interface{}
	array    []int
	count    int
	result   *PermResult
}

func (pt *Perm) perm(k int) {
	if k > pt.length {
		pt.count++
		pt.result = pt.f(pt.array[:len(pt.array)-1], pt.result, pt.argument)
		return
	}
	for i := k - 1; i >= 0; i-- {
		pt.array[i+1] = pt.array[i]
		pt.array[i] = k
		pt.perm(k + 1)
	}
	for i := 1; i < k; i++ {
		pt.array[i-1] = pt.array[i]
	}
}

// PermutationProcess will calc permutation, each permutation you can call function f
func PermutationProcess(length int, f func([]int, *PermResult, interface{}) *PermResult, argument interface{}) *Perm {
	pt := new(Perm)
	pt.argument = argument
	pt.length = length
	pt.array = make([]int, length+1)
	pt.count = 0
	pt.f = f
	pt.result = new(PermResult)
	pt.perm(1)
	return pt
}

// Sgn will return 1 when even permutation
// Sgn will return -1 when odd permutation
func Sgn(num []int) int {
	if len(num) <= 0 {
		return 0
	}
	result := 1
	length := len(num)
	if length < 2 {
		return 1
	}
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			result *= (num[i] - num[j])
			if result > 100000 {
				result /= int(math.Abs(float64(result)))
			}
		}
	}
	if result < 0 {
		return -1
	}
	return 1
}

func (m *Matrix) findNonZeroFromBelowRow(r, c int) (row int, err error) {
	if r > m.row || c > m.column {
		return -1, errors.New("The row or column are outside of this matrix")
	}
	for i := r + 1; i <= m.row; i++ {
		if val, _ := m.At(i, c); val != 0 {
			return i, nil
		}
	}
	return -1, errors.New("There is no available row")
}

// Determinant will calculate determinant
func (m *Matrix) Determinant() (float64, error) {
	if m.row != m.column {
		// TODO REUTN ERROR
		return 0, errors.New("This is not square error")
	}
	length := m.Row()
	matrix := Copy(m)
	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			if i < j {
				aji, _ := matrix.At(j, i)
				aii, _ := matrix.At(i, i)
				if aii != 0 {
					buf := aji / aii
					for k := 1; k <= length; k++ {
						ajk, _ := matrix.At(j, k)
						aik, _ := matrix.At(i, k)
						matrix.Set(j, k, ajk-aik*buf)
					}
				} else if aii == 0 {
					r, err := matrix.findNonZeroFromBelowRow(i, i)
					if err != nil {
						return 0, errors.New("there is no Determinant")
					}
					for k := 1; k <= length; k++ {
						aik, _ := matrix.At(i, k)
						ark, _ := matrix.At(r, k)
						matrix.Set(i, k, aik+ark)
						j--
					}
				}
			}
		}
	}
	result := float64(1)
	for i := 1; i <= length; i++ {
		aii, _ := matrix.At(i, i)
		result *= aii
	}
	return result, nil
}
