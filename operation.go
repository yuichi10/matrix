package matrix

import "errors"

// ZeroMatrix make all value 0
func (m *Matrix) ZeroMatrix() (matrix *Matrix) {
	matrix = Copy(m)
	matrix.matrix = make([]float64, m.row*m.column)
	return
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

func (m *Matrix) addColumnMatrix(mat Matrix) error {
	if m.row != mat.row {
		return errors.New("Row length is not same")
	}
	result := make([]float64, len(m.matrix)+len(mat.matrix))
	c := 0
	for i := 0; i < mat.row; i++ {
		for j := 0; j < m.column; j++ {
			result[c] = m.matrix[i*m.column+j]
			c++
		}
		for j := 0; j < mat.column; j++ {
			result[c] = mat.matrix[i*mat.column+j]
			c++
		}
	}
	m.column = m.column + mat.column
	m.matrix = result
	return nil
}

func (m *Matrix) addColumnArray(vector []float64) error {
	if m.row != len(vector) {
		return errors.New("Given vector and matrix row length is not same")
	}
	result := make([]float64, len(m.matrix)+len(vector))
	c := 0
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			result[c] = m.matrix[i*m.column+j]
			c++
		}
		result[c] = vector[i]
		c++
	}
	m.column++
	m.matrix = result
	return nil
}

func (m *Matrix) addColumnNumber(num float64) {
	result := make([]float64, len(m.matrix)+m.row)
	c := 0
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			result[c] = m.matrix[i*m.column+j]
			c++
		}
		result[c] = num
		c++
	}
	m.column++
	m.matrix = result
}

// AddColumn will add column
func (m *Matrix) AddColumn(num interface{}) (matrix *Matrix) {
	matrix = Copy(m)
	if mat, ok := num.(Matrix); ok {
		matrix.err = matrix.addColumnMatrix(mat)
		return
	} else if mat, ok := num.(*Matrix); ok {
		matrix.err = matrix.addColumnMatrix(*mat)
		return
	} else if row, ok := num.([]float64); ok {
		matrix.err = matrix.addColumnArray(row)
		return
	} else if row, ok := num.(int); ok {
		matrix.addColumnNumber(float64(row))
		return
	} else if row, ok := num.(float64); ok {
		matrix.addColumnNumber(row)
		return
	}
	matrix.err = errors.New("The argument type is not allowed")
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
