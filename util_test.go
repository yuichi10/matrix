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
