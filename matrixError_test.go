package matrix

import "testing"

func TestError(t *testing.T) {
	err := MatrixError{msg: "sample", funcName: "TestError"}
	ans := "[function: TestError] sample"
	if err.Error() != ans {
		t.Errorf("want %v but got %v", ans, err.Error())
	}
}
