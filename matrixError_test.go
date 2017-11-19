package matrix

import "testing"

func TestError(t *testing.T) {
	opt1 := New(3, 4, nil)
	err := MatrixError{msg: "sample", funcName: "TestError", opt1: opt1, opt2: nil}
	ans := "Size: opt1: (3, 4), opt2: No matrix\n[function: TestError] sample"
	if err.Error() != ans {
		t.Errorf("want \n %v but got \n %v", ans, err.Error())
	}
}
