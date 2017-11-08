package matrix

import (
	"errors"
)

func (m *Matrix) subByMatrix(mat Matrix) error {
	if err := m.checkSameSize(mat); err != nil {
		return err
	}
	for i, val := range mat.matrix {
		m.matrix[i] -= val
	}
	return nil
}

func (m *Matrix) addByMatrix(mat Matrix) error {
	if err := m.checkSameSize(mat); err != nil {
		return err
	}
	for i, val := range mat.matrix {
		m.matrix[i] += val
	}
	return nil
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
	if err := m.checkCanMulti(mat); err != nil {
		return err
	}
	matrix, _ := NewMatrix(m.row, mat.column)
	for i := 0; i < m.row*mat.column; i++ {
		matrix.matrix[i] = m.multiAtIndex(mat, i)
	}
	m.SetMatrix(matrix)
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
	if mat, ok := num.(Matrix); ok {
		return m.addByMatrix(mat)
	}
	if mat, ok := num.(*Matrix); ok {
		return m.addByMatrix(*mat)
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
	if mat, ok := num.(Matrix); ok {
		return m.subByMatrix(mat)
	} else if mat, ok := num.(*Matrix); ok {
		return m.subByMatrix(*mat)
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
		return m.multiByMatrix(mat)
	} else if mat, ok := num.(*Matrix); ok {
		return m.multiByMatrix(*mat)
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
