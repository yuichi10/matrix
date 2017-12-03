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

// Inverse will inverse matrix
func (m *Matrix) Inverse() (matrix *Matrix) {
	length := m.row
	matrix = NewEye(length)
	bufMatrix := Copy(m)
	if m.row != m.column {
		matrix.err = newError("The matrix is not square", "Inverse", m, nil)
		return
	}
	for i := 1; i <= length; i++ {
		val, _ := bufMatrix.At(i, i)
		if val <= 0.000001 && val >= -0.000001 {
			r, _ := bufMatrix.findNonZeroFromBelowRow(i, i)
			for n := 1; n <= length; n++ {
				brn, _ := bufMatrix.At(r, n)
				bin, _ := bufMatrix.At(i, n)
				bufMatrix.Set(i, n, bin+brn)
				mrn, _ := matrix.At(r, n)
				min, _ := matrix.At(i, n)
				matrix.Set(i, n, min+mrn)
			}
			val, _ = bufMatrix.At(i, i)
		}
		buf := 1 / val
		for j := 1; j <= length; j++ {
			bij, _ := bufMatrix.At(i, j)
			bufMatrix.Set(i, j, bij*buf)
			mij, _ := matrix.At(i, j)
			matrix.Set(i, j, mij*buf)
		}

		for k := 1; k <= length; k++ {
			if i != k {
				bki, _ := bufMatrix.At(k, i)

				for l := 1; l <= length; l++ {
					bkl, _ := bufMatrix.At(k, l)
					bil, _ := bufMatrix.At(i, l)
					bufMatrix.Set(k, l, bkl-bil*bki)
					mkl, _ := matrix.At(k, l)
					mil, _ := matrix.At(i, l)
					matrix.Set(k, l, mkl-mil*bki)
				}
			}
		}
	}
	return
}
