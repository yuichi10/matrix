package matrix

import "math"

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

// SigmoidGradient will return sigmoid gradient
func (m *Matrix) SigmoidGradient() (matrix *Matrix) {
	matrix = Copy(m)
	return matrix.Sigmoid().MultiEach(matrix.Sigmoid().MultiEach(-1).Add(1))
}

// Power will return power of matrix
func (m *Matrix) Power(e float64) (matrix *Matrix) {
	matrix = Copy(m)
	for i := 0; i < len(m.matrix); i++ {
		matrix.matrix[i] = math.Pow(matrix.matrix[i], e)
	}
	return
}
