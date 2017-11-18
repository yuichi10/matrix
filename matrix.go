package matrix

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Matrix has information of matrix
type Matrix struct {
	row    int       // 行
	column int       // 列
	matrix []float64 // 行 * 列の長さ
	err    error
}

func (m *Matrix) newByFloatArray(vector []float64) {
	if len(vector) == 0 {
		m.matrix = make([]float64, m.row*m.column)
		return
	}
	vec := make([]float64, len(vector))
	copy(vec, vector)
	m.matrix = vec
	if err := m.checkNormal(); err != nil {
		m.err = err
		return
	}
	return
}

func (m *Matrix) newByFLoat64(value float64) {
	m.matrix = make([]float64, m.row*m.column)
	for i := 0; i < m.row*m.column; i++ {
		m.matrix[i] = value
	}
}

// New will return *Matrix
func New(row, column int, value interface{}) (matrix *Matrix) {
	matrix = new(Matrix)
	matrix.row = row
	matrix.column = column
	if row <= 0 || column <= 0 {
		matrix.err = errors.New("Length is not greater 0")
		return
	}
	if vector, ok := value.([]float64); ok {
		matrix.newByFloatArray(vector)
		return
	} else if num, ok := value.(int); ok {
		matrix.newByFLoat64(float64(num))
		return
	} else if num, ok := value.(float64); ok {
		matrix.newByFLoat64(num)
		return
	} else if value == nil {
		matrix.matrix = make([]float64, row*column)
		return
	}
	matrix.err = errors.New("The argument type is not allowed")
	return
}

// NewVector will create vector by array
func NewVector(row []float64) (matrix *Matrix) {
	matrix = new(Matrix)
	if len(row) <= 0 {
		matrix.err = errors.New("The vector is broken")
		return
	}
	vector := make([]float64, len(row))
	copy(vector, row)
	matrix.row = len(row)
	matrix.column = 1
	matrix.matrix = vector
	return
}

// NewRandom will return matrix which values are 0~1
func NewRandom(row, column int, digits uint8) (matrix *Matrix) {
	matrix = new(Matrix)
	if row <= 0 || column <= 0 {
		matrix.err = errors.New("Length is not greater 0")
		return
	}
	matrix.row = row
	matrix.column = column
	matrix.matrix = make([]float64, row*column)
	d := math.Pow10(int(digits))
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	for i := 0; i < row*column; i++ {
		matrix.matrix[i] = float64(rand.Intn(int(d))) / d
	}
	return
}

// NewHotVector will return hot vector
func NewHotVector(size, place int) (matrix *Matrix) {
	matrix = new(Matrix)
	if size <= 0 || place <= 0 {
		matrix.err = errors.New("The size and place must be > 0")
		return
	} else if place > size {
		matrix.err = errors.New("place must be less or equal than size")
		return
	}
	matrix.row = size
	matrix.column = 1
	matrix.matrix = make([]float64, size)
	matrix.matrix[place-1] = 1
	return
}

// Err will return error of calcuration
func (m *Matrix) Err() error {
	return m.err
}

// Copy will copy matrix
func Copy(mat *Matrix) *Matrix {
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	matrix := &Matrix{mat.row, mat.column, vector, mat.err}
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
func (m *Matrix) AddRow(num interface{}) (matrix *Matrix) {
	matrix = Copy(m)
	if mat, ok := num.(Matrix); ok {
		matrix.err = matrix.addRowMatrix(mat)
		return
	} else if mat, ok := num.(*Matrix); ok {
		matrix.err = matrix.addRowMatrix(*mat)
		return
	} else if row, ok := num.([]float64); ok {
		if matrix.column != len(row) {
			matrix.err = errors.New("Column length is not same")
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
	matrix.err = errors.New("The argument type is not allowed")
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
func (m *Matrix) AddRowHEAD(num interface{}) (matrix *Matrix) {
	matrix = Copy(m)
	if mat, ok := num.(Matrix); ok {
		matrix.err = matrix.addRowMatrixHEAD(mat)
		return
	} else if mat, ok := num.(*Matrix); ok {
		matrix.err = matrix.addRowMatrixHEAD(*mat)
		return
	} else if row, ok := num.([]float64); ok {
		if matrix.column != len(row) {
			matrix.err = errors.New("Column length is not same")
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
	matrix.err = errors.New("The argument type is not allowed")
	return
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

// Reshape will change vect
func (m *Matrix) Reshape(row, column int) (matrix *Matrix) {
	matrix = Copy(m)
	matrix.row = row
	matrix.column = column
	if err := matrix.checkNormal(); err != nil {
		matrix.err = err
		return
	}
	return
}

// SepRow will return matrix which separate by row numbers
func (m *Matrix) SepRow(start, end int) (matrix *Matrix) {
	matrix = Copy(m)
	if end < start {
		matrix.err = errors.New("The argument values are invalid")
		return
	} else if end > m.row || start < 1 {
		matrix.err = errors.New("The value are out of matrix")
		return
	}
	s := (start - 1) * m.column
	e := (end - 1) * m.column
	matrix = New(end-start+1, m.column, m.matrix[s:e+m.column])
	return
}

// SepColumn will return matrix which separate by sep numbers
func (m *Matrix) SepColumn(start, end int) (matrix *Matrix) {
	matrix = Copy(m)
	if end < start {
		matrix.err = errors.New("The argument values are invalid")
		return
	} else if end > m.column || start < 1 {
		matrix.err = errors.New("The value are out of matrix")
		return
	}
	vector := make([]float64, (end-start+1)*m.row)
	count := 0
	for i := 0; i < m.row; i++ {
		for j := start - 1; j < end; j++ {
			vector[count] = m.matrix[i*m.column+j]
			count++
		}
	}
	matrix = New(m.row, end-start+1, vector)
	return
}
