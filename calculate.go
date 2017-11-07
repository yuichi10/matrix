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

// index will start from 1
func (m *Matrix) multiAtIndex(mat Matrix, index int) float64 {
	var val float64
	r := (index) / mat.column
	c := (index) % mat.column
	for i := 0; i < m.column; i++ {
		val += m.matrix[i+r*m.column] * mat.matrix[i*mat.column+c]
	}
	return val
}

func (m *Matrix) multiByMatrix(mat Matrix) error {
	matrix, err := NewMatrix(m.row, mat.column)
	if err != nil {
		return err
	}
	for i := 0; i < m.row*mat.column; i++ {
		matrix.matrix[i] = m.multiAtIndex(mat, i)
	}
	if err := m.SetMatrix(*matrix); err != nil {
		return err
	}
	return nil
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

func (m *Matrix) multiByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] *= num
	}
}

// Add will add some value to Matrix
func (m *Matrix) Add(num interface{}) error {
	m.checkNormal()
	if mat, ok := num.(Matrix); ok {
		if mat.row != m.row || mat.column != m.column {
			return errors.New("The row and column num are different")
		}
		m.addByMatrix(mat)
		return nil
	}
	if mat, ok := num.(*Matrix); ok {
		if mat.row != m.row || mat.column != m.column {
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

// Sub will calculate sub of matrix
func (m *Matrix) Sub(num interface{}) error {
	if err := m.checkNormal(); err != nil {
		return err
	}

	if mat, ok := num.(Matrix); ok {
		if err := m.checkSameSize(mat); err != nil {
			return err
		}
		m.subByMatrix(mat)
		return nil
	} else if mat, ok := num.(*Matrix); ok {
		if err := m.checkSameSize(*mat); err != nil {
			return err
		}
		m.subByMatrix(*mat)
		return nil
	} else if mat, ok := num.(float64); ok {
		m.subByFloat(mat)
		return nil
	} else if mat, ok := num.(int); ok {
		m.subByFloat(float64(mat))
		return nil
	}
	return errors.New("The sub op2 type is not allowed")
}

// Multi will calculate Multi
func (m *Matrix) Multi(num interface{}) error {
	if mat, ok := num.(Matrix); ok {
		if err := m.checkCanMulti(mat); err != nil {
			return err
		}
		m.multiByMatrix(mat)
		return nil
	} else if mat, ok := num.(*Matrix); ok {
		if err := m.checkCanMulti(*mat); err != nil {
			return err
		}
		m.multiByMatrix(*mat)
		return nil
	} else if mat, ok := num.(float64); ok {
		m.multiByFloat(float64(mat))
		return nil
	} else if mat, ok := num.(int); ok {
		m.multiByFloat(float64(mat))
		return nil
	}
	return errors.New("The multi op2 type is not allowed")
}

// MultiEach
// Sep
