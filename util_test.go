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
