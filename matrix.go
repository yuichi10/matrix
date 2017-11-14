package matrix

import (
	"errors"
	"fmt"
	"strings"
)

// Matrix has information of matrix
type Matrix struct {
	row     int       // 行
	column  int       // 列
	matrix  []float64 // 行 * 列の長さ
	calcErr error
}

// New will return *Matrix
func New(row, column int, vector []float64) (*Matrix, error) {
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

// CalcErr will return error of calcuration
func (m *Matrix) CalcErr() error {
	return m.calcErr
}

// Copy will copy matrix
func Copy(mat *Matrix) *Matrix {
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	matrix := &Matrix{mat.row, mat.column, vector, mat.calcErr}
	return matrix
}

// AddRowMatrix will add matrix behinde this matrix
func (m *Matrix) addRowMatrix(mat Matrix) error {
	if m.column != mat.column {
		return errors.New("Column length is not same")
	}
	m.matrix = append(m.matrix, mat.matrix...)
	m.row += mat.row
	return nil
}

// AddRow add row at tail. if the len of column = 0. create new vector 1 * len(row)
// []float64, Matrix, int and float64 are only allowed
func (m *Matrix) AddRow(num interface{}) (matrix *Matrix, err error) {
	matrix = Copy(m)
	if mat, ok := num.(Matrix); ok {
		err = matrix.addRowMatrix(mat)
		return
	} else if mat, ok := num.(*Matrix); ok {
		err = matrix.addRowMatrix(*mat)
		return
	} else if row, ok := num.([]float64); ok {
		if matrix.column != len(row) {
			err = errors.New("Column length is not same")
			return
		}
		matrix.row++
		matrix.matrix = append(matrix.matrix, row...)
		return
	} else if row, ok := num.(int); ok {
		matrix.row++
		vector := make([]float64, matrix.column)
		for i := range vector {
			vector[i] = float64(row)
		}
		matrix.matrix = append(matrix.matrix, vector...)
		return
	} else if row, ok := num.(float64); ok {
		matrix.row++
		vector := make([]float64, matrix.column)
		for i := range vector {
			vector[i] = float64(row)
		}
		matrix.matrix = append(matrix.matrix, vector...)
		return
	}
	err = errors.New("The argument type is not allowed")
	return
}

// AddRowMatrixHEAD will add matrix HEAD this matrix
func (m *Matrix) addRowMatrixHEAD(mat Matrix) error {
	if m.column != mat.column {
		return errors.New("Column length is not same")
	}
	m.matrix = append(mat.matrix, m.matrix...)
	m.row += mat.row
	return nil
}

// AddRowHEAD add row at head. if the len of column = 0 create new vector
func (m *Matrix) AddRowHEAD(num interface{}) (matrix *Matrix, err error) {
	matrix = Copy(m)
	if mat, ok := num.(Matrix); ok {
		err = matrix.addRowMatrixHEAD(mat)
		return
	} else if mat, ok := num.(*Matrix); ok {
		err = matrix.addRowMatrixHEAD(*mat)
		return
	} else if row, ok := num.([]float64); ok {
		if matrix.column != len(row) {
			err = errors.New("Column length is not same")
			return
		}
		matrix.row++
		matrix.matrix = append(row, matrix.matrix...)
		return
	} else if row, ok := num.(int); ok {
		matrix.row++
		vector := make([]float64, matrix.column)
		for i := range vector {
			vector[i] = float64(row)
		}
		matrix.matrix = append(vector, matrix.matrix...)
		return
	} else if row, ok := num.(float64); ok {
		matrix.row++
		vector := make([]float64, matrix.column)
		for i := range vector {
			vector[i] = float64(row)
		}
		matrix.matrix = append(vector, matrix.matrix...)
		return
	}
	err = errors.New("The argument type is not allowed")
	return
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

// Row return this matrix's row
func (m *Matrix) Row() int {
	return m.row
}

// Column return this matrix's column
func (m *Matrix) Column() int {
	return m.column
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

// SepRow will return matrix which separate by row numbers
func (m *Matrix) SepRow(start, end int) (*Matrix, error) {
	if end < start {
		return nil, errors.New("The argument values are invalid")
	} else if end > m.row || start < 1 {
		return nil, errors.New("The value are out of matrix")
	}
	s := (start - 1) * m.column
	e := (end - 1) * m.column
	matrix, err := New(end-start+1, m.column, m.matrix[s:e+m.column])
	return matrix, err
}

// SepColumn will return matrix which separate by sep numbers
func (m *Matrix) SepColumn(start, end int) (*Matrix, error) {
	if end < start {
		return nil, errors.New("The argument values are invalid")
	} else if end > m.column || start < 1 {
		return nil, errors.New("The value are out of matrix")
	}
	vector := make([]float64, (end-start+1)*m.row)
	count := 0
	for i := 0; i < m.row; i++ {
		for j := start - 1; j < end; j++ {
			vector[count] = m.matrix[i*m.column+j]
			count++
		}
	}
	return New(m.row, end-start+1, vector)
}
