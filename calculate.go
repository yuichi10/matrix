package matrix

import (
	"errors"
)

func (m *Matrix) subByMatrix(mat Matrix) {
	for i, val := range mat.matrix {
		m.matrix[i] -= val
	}
}

func (m *Matrix) addByMatrix(mat Matrix) {
	for i, val := range mat.matrix {
		m.matrix[i] += val
	}
}

func (m *Matrix) subByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] -= num
	}
}

func (m *Matrix) addByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] += num
	}
}

// Add will add some value to Matrix
func (m *Matrix) Add(num interface{}) error {
	m.isNormal()
	if mat, ok := num.(Matrix); ok {
		if mat.rows != m.rows || mat.columns != m.columns {
			return errors.New("The row and column num are different")
		}
		m.addByMatrix(mat)
		return nil
	}
	if mat, ok := num.(*Matrix); ok {
		if mat.rows != m.rows || mat.columns != m.columns {
			return errors.New("The row and column num are different")
		}
		m.addByMatrix(*mat)
		return nil
	}
	if mat, ok := num.(int); ok {
		m.addByFloat(float64(mat))
		return nil
	}
	if mat, ok := num.(float64); ok {
		m.addByFloat(float64(mat))
		return nil
	}
	return errors.New("The add type is not allowed")
}

// Sub

// Multi
// MultiEach
// Sep
