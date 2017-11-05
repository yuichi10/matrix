package matrix

import (
	"errors"
	"log"
)

// checkBroken look the length of row and column if t
func (m *Matrix) isNormal() (bool, error) {
	if m.rows <= 0 || m.columns <= 0 {
		log.Fatal("Matrix is broken")
		return false, errors.New("Matrix size is broken")
	}
	if len(m.matrix) != m.rows*m.columns {
		return false, errors.New("matrix size and row/colmn relationship is broken")
	}
	return true, nil
}

// checkSize will check argument row and column is in size
func (m *Matrix) isThereValue(row, column int) bool {
	if row <= 0 || column <= 0 || row > m.rows || column > m.columns {
		log.Fatal("Size is invalid")
		return false
	}
	return true
}

func (m *Matrix) isSameSize(mat Matrix) bool {
	if mat.rows != m.rows || mat.columns != m.columns {
		return false
	}
	return true
}
