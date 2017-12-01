package matrix

import "testing"

func TestError(t *testing.T) {
	var opt1 *Matrix
	var opt2 *Matrix
	var ans string
	var err *matrixError
	opt1 = New(3, 4, nil)
	err = newError("sample", "TestError", opt1, nil)
	ans = "Size: opt1: (3, 4), opt2: No matrix\n[function: TestError] sample"
	if err.Error() != ans {
		t.Errorf("want \n %v but got \n %v", ans, err.Error())
	}

	opt1 = New(3, 4, nil)
	opt2 = New(4, 5, nil)
	err = newError("sample", "TestError", opt1, opt2)
	ans = "Size: opt1: (3, 4), opt2: (4, 5)\n[function: TestError] sample"
	if err.Error() != ans {
		t.Errorf("want \n %v \nbut got \n %v", ans, err.Error())
	}

	Config.Panic = true
	opt1 = New(3, 4, nil)
	opt2 = New(4, 5, nil)
	f := func() { newError("sample", "TestError", opt1, opt2) }
	assertPanic(t, f)
	Config.Panic = false
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
