package matrix

import "math"

// Transpose will make transposed vector
func (m *Matrix) Transpose() (matrix *Matrix) {
	matrix = Copy(m)
	count := 0
	for i := 0; i < matrix.column; i++ {
		for j := 0; j < matrix.row; j++ {
			matrix.matrix[count] = m.matrix[j*m.column+i]
			count++
		}
	}
	matrix.row = m.column
	matrix.column = m.row
	return
}

// Vector will return vector version of this matrix
func (m *Matrix) Vector() (matrix *Matrix) {
	matrix = Copy(m)
	matrix.row = m.row * m.column
	matrix.column = 1
	return
}

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() (matrix *Matrix) {
	matrix = Copy(m)
	matrix.matrix = make([]float64, m.row*m.column)
	return
}

// Sigmoid will return sigmoied matrix
func (m *Matrix) Sigmoid() (matrix *Matrix) {
	matrix = Copy(m)
	for i := 1; i <= matrix.row; i++ {
		for j := 1; j <= matrix.column; j++ {
			val, _ := matrix.At(i, j)
			sig := 1.0 / (1.0 + math.Exp(-val))
			matrix.Set(i, j, sig)
		}
	}
	return
}
