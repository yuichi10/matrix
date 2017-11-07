package matrix

import (
	"errors"
)

// checkBroken look the length of row and column if t
func (m *Matrix) checkNormal() error {
	if m.rows <= 0 || m.columns <= 0 {
		return errors.New("Matrix size is broken")
	}
	if len(m.matrix) != m.rows*m.columns {
		return errors.New("matrix size and row/colmn relationship is broken")
	}
	return nil
}

// checkSize will check argument row and column is in size
func (m *Matrix) checkThereValue(row, column int) error {
	if row <= 0 || column <= 0 || row > m.rows || column > m.columns {
		return errors.New("There is out of matrix")
	}
	return nil
}

func (m *Matrix) checkSameSize(mat Matrix) error {
	if mat.rows != m.rows || mat.columns != m.columns {
		return errors.New("The size is not same")
	}
	return nil
}
