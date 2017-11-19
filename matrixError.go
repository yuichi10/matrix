package matrix

import "fmt"

// MatrixError is struct of error
// to make easy to show error message
type MatrixError struct {
	msg      string
	funcName string
	opt1     *Matrix
	opt2     *Matrix
}

func (err *MatrixError) sizeDetail() string {
	opt1 := "opt1: No matrix"
	opt2 := "opt2: No matrix"
	if err.opt1 != nil {
		r, c := err.opt1.Size()
		opt1 = fmt.Sprintf("opt1: (%v, %v)", r, c)
	}
	if err.opt2 != nil {
		r, c := err.opt2.Size()
		opt2 = fmt.Sprintf("opt2: (%v, %v)", r, c)
	}
	return fmt.Sprintf("Size: %v, %v", opt1, opt2)
}

// MyError構造体にerrorインタフェースのError関数を実装
func (err *MatrixError) Error() string {
	size := err.sizeDetail()
	return fmt.Sprintf("%v\n[function: %s] %s", size, err.funcName, err.msg)
}
