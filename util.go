package matrix

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
