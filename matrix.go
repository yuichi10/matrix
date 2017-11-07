package matrix

import (
	"errors"
	"fmt"
)

// Matrix has information of matrix
type Matrix struct {
	rows    int       // 行
	columns int       // 列
	matrix  []float64 // 行 * 列の長さ
}

// NewMatrix will return *Matrix
func NewMatrix(row, column int) (*Matrix, error) {
	matrix := new(Matrix)
	if row <= 0 || column <= 0 {
		return nil, errors.New("Length is not greater equal 0")
	}
	matrix.rows = row
	matrix.columns = column
	matrix.matrix = make([]float64, matrix.rows*matrix.columns)
	return matrix, nil
}

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() {
	m.checkNormal()
	m.matrix = make([]float64, m.rows*m.columns)
}

// MakeVector will create vector by array
func (m *Matrix) MakeVector(row []float64) {
	m.rows = 1
	m.columns = len(row)
	m.matrix = row
}

// AddRow add row at tail. if the len of column = 0. create new vector 1 * len(row)
func (m *Matrix) AddRow(row []float64) error {
	if m.columns != len(row) && m.columns != 0 {
		return errors.New("Column length is not same")
	}
	if m.columns == 0 {
		m.MakeVector(row)
		return nil
	}
	m.rows++
	m.matrix = append(m.matrix, row...)
	return nil
}

// AddRowHEAD add row at head. if the len of column = 0 create new vector
func (m *Matrix) AddRowHEAD(row []float64) error {
	if m.columns != len(row) && m.columns != 0 {
		return errors.New("Column length is not same")
	}
	if m.columns == 0 {
		m.MakeVector(row)
		return nil
	}
	m.rows++
	m.matrix = append(row, m.matrix...)
	return nil
}

// Show will show matrix condition
func (m *Matrix) Show() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			fmt.Print(m.matrix[i*m.columns+j])
		}
		fmt.Println()
	}
}

// Size return matrix size
func (m *Matrix) Size() (int, int) {
	return m.rows, m.columns
}

// At return a point of value
func (m *Matrix) At(row, column int) (float64, error) {
	if err := m.checkThereValue(row, column); err != nil {
		return 0, err
	}
	return m.matrix[column*(row-1)+column-1], nil
}

// Set will set specifix value
func (m *Matrix) Set(row, column int, value float64) error {
	if err := m.checkThereValue(row, column); err != nil {
		return err
	}
	m.matrix[column*(row-1)+column-1] = value
	return nil
}
