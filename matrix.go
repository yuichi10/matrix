package matrix

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yuichi10/matrix/config"
)

// Matrix has information of matrix
type Matrix struct {
	row    int       // 行
	column int       // 列
	matrix []float64 // 行 * 列の長さ
	err    error
}

var (
	// Config has configure of this matrix library
	Config *config.Config
)

func init() {
	Config = new(config.Config)
	Config.Panic = false
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
	text := ""
	for i := 0; i < m.row; i++ {
		line := ""
		for j := 0; j < m.column; j++ {
			line = fmt.Sprintf("%v %.5f", line, m.matrix[i*m.column+j])
		}
		text += strings.Trim(line, " ") + "\n"
	}
	if m.err != nil {
		text += m.err.Error() + "\n"
	}
	fmt.Println(text)
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

// Int will make all value int
func (m *Matrix) Int() (matrix *Matrix) {
	matrix = Copy(m)
	for i := 0; i < len(matrix.matrix); i++ {
		matrix.matrix[i] = float64(int(matrix.matrix[i]))
	}
	return
}
