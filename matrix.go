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
	err    error
}

// Row return this matrix's row
func (m *Matrix) Row() int {
	return m.row
}

// Column return this matrix's column
func (m *Matrix) Column() int {
	return m.column
}

// Size return matrix size
func (m *Matrix) Size() (int, int) {
	return m.row, m.column
}

// Show will show matrix condition
func (m *Matrix) Show() {
	fmt.Printf("size: %v x %v\n", m.Row(), m.Column())
	for i := 0; i < m.row; i++ {
		line := ""
		for j := 0; j < m.column; j++ {
			line = fmt.Sprintf("%v %v", line, m.matrix[i*m.column+j])
		}
		fmt.Println(strings.Trim(line, " "))
	}
}

// Err will return error of calcuration
func (m *Matrix) Err() error {
	return m.err
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
func (m *Matrix) SetMatrix(mat *Matrix) {
	if err := mat.checkNormal(); err != nil {
		m.err = errors.New("The matrix is broken")
		return
	}
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	m.row = mat.row
	m.column = mat.column
	m.matrix = vector
	return
}
