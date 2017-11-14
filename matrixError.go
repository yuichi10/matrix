package matrix

import "fmt"

// MatrixError is struct of error
// to make easy to show error message
type MatrixError struct {
	msg      string
	funcName string
}

// MyError構造体にerrorインタフェースのError関数を実装
func (err *MatrixError) Error() string {
	return fmt.Sprintf("[function: %s] %s", err.funcName, err.msg)
}
