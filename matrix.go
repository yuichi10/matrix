package matrix

import (
	"errors"
	"fmt"
)

// Matrix has information of matrix
type Matrix struct {
	row    int       // 行
	column int       // 列
	matrix []float64 // 行 * 列の長さ
}

// NewMatrix will return *Matrix
func NewMatrix(row, column int) (*Matrix, error) {
	matrix := new(Matrix)
	if row <= 0 || column <= 0 {
		return nil, errors.New("Length is not greater equal 0")
	}
	matrix.row = row
	matrix.column = column
	matrix.matrix = make([]float64, matrix.row*matrix.column)
	return matrix, nil
}

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() {
	m.checkNormal()
	m.matrix = make([]float64, m.row*m.column)
}

// MakeVector will create vector by array
func (m *Matrix) MakeVector(row []float64) {
	// TODO: check the vector
	m.row = 1
	m.column = len(row)
	m.matrix = row
}

// AddRow add row at tail. if the len of column = 0. create new vector 1 * len(row)
func (m *Matrix) AddRow(row []float64) error {
	if m.column != len(row) && m.column != 0 {
		return errors.New("Column length is not same")
	}
	if m.column == 0 {
		m.MakeVector(row)
		return nil
	}
	m.row++
	m.matrix = append(m.matrix, row...)
	return nil
}

// AddRowHEAD add row at head. if the len of column = 0 create new vector
func (m *Matrix) AddRowHEAD(row []float64) error {
	if m.column != len(row) && m.column != 0 {
		return errors.New("Column length is not same")
	}
	if m.column == 0 {
		m.MakeVector(row)
		return nil
	}
	m.row++
	m.matrix = append(row, m.matrix...)
	return nil
}

// Show will show matrix condition
func (m *Matrix) Show() {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			fmt.Print(m.matrix[i*m.column+j])
		}
		fmt.Println()
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
func (m *Matrix) SetMatrix(mat Matrix) error {
	if err := mat.checkNormal(); err != nil {
		return errors.New("The matrix is broken")
	}
	m.row = mat.row
	m.column = mat.column
	m.matrix = mat.matrix
	return nil
}
