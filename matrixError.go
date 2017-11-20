package matrix

import "fmt"

// matrixError is struct of error
// to make easy to show error message
type matrixError struct {
	msg      string
	funcName string
	opt1     *Matrix
	opt2     *Matrix
}

func newError(msg, funcName string, opt1, opt2 *Matrix) *matrixError {
	err := &matrixError{
		msg:      msg,
		funcName: funcName,
		opt1:     opt1,
		opt2:     opt2,
	}
	if Config.Panic {
		panic(err.Error())
	}
	return err
}

func (err *matrixError) sizeDetail() string {
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
func (err *matrixError) Error() string {
	size := err.sizeDetail()
	return fmt.Sprintf("%v\n[function: %s] %s", size, err.funcName, err.msg)
}
