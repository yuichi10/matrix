package matrix

import (
	"errors"
	"fmt"
	"strings"
)

// Matrix has information of matrix
type Matrix struct {
	row    int       // 行
	column int       // 列
	matrix []float64 // 行 * 列の長さ
}

// NewMatrix will return *Matrix
func NewMatrix(row, column int, vector []float64) (*Matrix, error) {
	matrix := new(Matrix)
	if row <= 0 || column <= 0 {
		return nil, errors.New("Length is not greater 0")
	}
	matrix.row = row
	matrix.column = column
	if len(vector) == 0 || vector == nil {
		matrix.matrix = make([]float64, matrix.row*matrix.column)
		return matrix, nil
	}
	vec := make([]float64, len(vector))
	copy(vec, vector)
	matrix.matrix = vec
	if err := matrix.checkNormal(); err != nil {
		return nil, err
	}
	return matrix, nil
}

// NewVector will create vector by array
func NewVector(row []float64) (*Matrix, error) {
	if len(row) <= 0 {
		return nil, errors.New("The vector is broken")
	}
	matrix := new(Matrix)
	vector := make([]float64, len(row))
	copy(vector, row)
	matrix.row = len(row)
	matrix.column = 1
	matrix.matrix = vector
	return matrix, nil
}

// Copy will copy matrix
func Copy(mat *Matrix) (*Matrix, error) {
	if err := mat.checkNormal(); err != nil {
		return nil, err
	}
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	matrix := &Matrix{mat.row, mat.column, vector}
	return matrix, nil
}

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() {
	m.matrix = make([]float64, m.row*m.column)
}

// AddRowMatrix will add matrix behinde this matrix
func (m *Matrix) AddRowMatrix(mat Matrix) error {
	if m.column != mat.column {
		return errors.New("Column length is not same")
	}
	m.matrix = append(m.matrix, mat.matrix...)
	m.row += mat.row
	return nil
}

// AddRow add row at tail. if the len of column = 0. create new vector 1 * len(row)
// float64 and Matrix are only allowed
func (m *Matrix) AddRow(num interface{}) error {
	if mat, ok := num.(Matrix); ok {
		return m.AddRowMatrix(mat)
	} else if mat, ok := num.(*Matrix); ok {
		return m.AddRowMatrix(*mat)
	} else if row, ok := num.([]float64); ok {
		if m.column != len(row) {
			return errors.New("Column length is not same")
		}
		m.row++
		m.matrix = append(m.matrix, row...)
		return nil
	} else if row, ok := num.(int); ok {
		m.row++
		vector := make([]float64, m.column)
		for i := range vector {
			vector[i] = float64(row)
		}
		m.matrix = append(m.matrix, vector...)
		return nil
	}
	return errors.New("The argument type is not allowed")
}

// AddRowHEAD add row at head. if the len of column = 0 create new vector
func (m *Matrix) AddRowHEAD(row []float64) error {
	if m.column != len(row) {
		return errors.New("Column length is not same")
	}
	m.row++
	m.matrix = append(row, m.matrix...)
	return nil
}

// Show will show matrix condition
func (m *Matrix) Show() {
	for i := 0; i < m.row; i++ {
		line := ""
		for j := 0; j < m.column; j++ {
			line = fmt.Sprintf("%v %v", line, m.matrix[i*m.column+j])
		}
		fmt.Println(strings.Trim(line, " "))
	}
}

// Size return matrix size
func (m *Matrix) Size() (int, int) {
	return m.row, m.column
}

// At return a point of value
func (m *Matrix) At(row, column int) (float64, error) {
	if err := m.checkThereValue(row, column); err != nil {
		return 0, err
	}
	return m.matrix[m.column*(row-1)+column-1], nil
}

// Set will set specifix value
func (m *Matrix) Set(row, column int, value float64) error {
	if err := m.checkThereValue(row, column); err != nil {
		return err
	}
	m.matrix[m.column*(row-1)+column-1] = value
	return nil
}

// SetMatrix will set mat to this matrix
func (m *Matrix) SetMatrix(mat *Matrix) error {
	if err := mat.checkNormal(); err != nil {
		return errors.New("The matrix is broken")
	}
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	m.row = mat.row
	m.column = mat.column
	m.matrix = vector
	return nil
}

// Transpose will make transposed vector
func (m *Matrix) Transpose() {
	vector := make([]float64, len(m.matrix))
	copy(vector, m.matrix)
	count := 0
	for i := 0; i < m.column; i++ {
		for j := 0; j < m.row; j++ {
			vector[count] = m.matrix[j*m.column+i]
			count++
		}
	}
	r := m.row
	c := m.column
	m.row = c
	m.column = r
	m.matrix = vector
}
