package matrix

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
	"time"
)

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
func NewRandom(row, column int, decimal uint8) (matrix *Matrix) {
	matrix = new(Matrix)
	if row <= 0 || column <= 0 {
		matrix.err = errors.New("Length is not greater 0")
		return
	}
	matrix.row = row
	matrix.column = column
	matrix.matrix = make([]float64, row*column)
	d := math.Pow10(int(decimal))
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

// NewEye will return Unit matrix
func NewEye(length int) (matrix *Matrix) {
	matrix = new(Matrix)
	if length <= 0 {
		matrix.err = newError("lenght should greater than 0", "NewEye", matrix, nil)
		return
	}
	matrix.row = length
	matrix.column = length
	matrix.matrix = make([]float64, length*length)
	for i := 0; i < length; i++ {
		matrix.matrix[matrix.column*i+i] = 1
	}
	return
}

// Copy will copy matrix
func Copy(mat *Matrix) *Matrix {
	vector := make([]float64, len(mat.matrix))
	copy(vector, mat.matrix)
	matrix := &Matrix{mat.row, mat.column, vector, mat.err}
	return matrix
}
