package matrix

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
