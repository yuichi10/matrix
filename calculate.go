package matrix

import (
	"context"
	"sync"
)

type calcMultiInfo struct {
	result *Matrix
	op2    *Matrix
	index  int
}

var wg sync.WaitGroup

func (m *Matrix) subByMatrix(mat Matrix) error {
	if err := m.checkSameSize(mat); err != nil {
		return err
	}
	for i, val := range mat.matrix {
		m.matrix[i] -= val
	}
	return nil
}

func (m *Matrix) addByMatrix(mat Matrix) error {
	if err := m.checkSameSize(mat); err != nil {
		return err
	}
	for i, val := range mat.matrix {
		m.matrix[i] += val
	}
	return nil
}

// index will start from 1
func (m *Matrix) multiAtIndex(mat Matrix, index int) float64 {
	var val float64
	r := (index) / mat.column
	c := (index) % mat.column
	for i := 0; i < m.column; i++ {
		val += m.matrix[i+r*m.column] * mat.matrix[i*mat.column+c]
	}
	return val
}

func (m *Matrix) multiAtIndexParallel(ctx context.Context, calcInfo chan *calcMultiInfo) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case info := <-calcInfo:
			var val float64
			r := (info.index) / info.op2.column
			c := (info.index) % info.op2.column
			for i := 0; i < m.column; i++ {
				val += m.matrix[i+r*m.column] * info.op2.matrix[i*info.op2.column+c]
			}
			info.result.matrix[info.index] = val
		}
	}
}

func (m *Matrix) multiByMatrix(mat Matrix) error {
	if err := m.checkCanMulti(mat); err != nil {
		return err
	}
	matrix, _ := New(m.row, mat.column, nil)
	for i := 0; i < m.row*mat.column; i++ {
		matrix.matrix[i] = m.multiAtIndex(mat, i)
	}
	m.SetMatrix(matrix)
	return nil
}

func (m *Matrix) multiByMatrixParallel(mat Matrix) error {
	if err := m.checkCanMulti(mat); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	calcInfo := make(chan *calcMultiInfo)
	wg.Add(1)
	go m.multiAtIndexParallel(ctx, calcInfo)
	matrix, _ := New(m.row, mat.column, nil)
	for i := 0; i < m.row*mat.column; i++ {
		info := &calcMultiInfo{result: matrix, op2: &mat, index: i}
		calcInfo <- info
	}
	cancel()
	wg.Wait()
	m.SetMatrix(matrix)
	return nil
}

func (m *Matrix) multiEachByMatrix(mat Matrix) error {
	if err := m.checkSameSize(mat); err != nil {
		return err
	}
	for i, val := range mat.matrix {
		m.matrix[i] *= val
	}
	return nil
}

func (m *Matrix) subByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] -= num
	}
}

func (m *Matrix) addByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] += num
	}
}

func (m *Matrix) multiByFloat(num float64) {
	for i := range m.matrix {
		m.matrix[i] *= num
	}
}

// Add will add some value to Matrix
func (m *Matrix) Add(num interface{}) (matrix *Matrix) {
	var err error
	matrix = Copy(m)
	if m.CalcErr() != nil {
		return
	}
	if mat, ok := num.(Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.addByMatrix(mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Add"}
		}
		return
	}
	if mat, ok := num.(*Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.addByMatrix(*mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Add"}
		}
		return
	}
	if mat, ok := num.(int); ok {
		matrix.addByFloat(float64(mat))
		return
	}
	if mat, ok := num.(float64); ok {
		matrix.addByFloat(float64(mat))
		return
	}
	matrix.calcErr = &MatrixError{msg: "The add type is not allowed", funcName: "Add"}
	return
}

// Sub will calculate sub of matrix
func (m *Matrix) Sub(num interface{}) (matrix *Matrix) {
	var err error
	matrix = Copy(m)
	if m.CalcErr() != nil {
		return
	}
	if mat, ok := num.(Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.subByMatrix(mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Sub"}
		}
		return
	} else if mat, ok := num.(*Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.subByMatrix(*mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Sub"}
		}
		return
	} else if mat, ok := num.(float64); ok {
		matrix.subByFloat(mat)
		return
	} else if mat, ok := num.(int); ok {
		matrix.subByFloat(float64(mat))
		return
	}
	matrix.calcErr = &MatrixError{msg: "The sub op2 type is not allowed", funcName: "Sub"}
	return
}

// Multi will calculate Multi
func (m *Matrix) Multi(num interface{}) (matrix *Matrix) {
	var err error
	matrix = Copy(m)
	if m.CalcErr() != nil {
		return
	}
	if mat, ok := num.(Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.multiByMatrixParallel(mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Multi"}
		}
		return
	} else if mat, ok := num.(*Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.multiByMatrixParallel(*mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "Multi"}
		}
		return
	} else if mat, ok := num.(float64); ok {
		matrix.multiByFloat(float64(mat))
		return
	} else if mat, ok := num.(int); ok {
		matrix.multiByFloat(float64(mat))
		return
	}
	matrix.calcErr = &MatrixError{msg: "The multi op2 type is not allowed", funcName: "Multi"}
	return
}

// MultiEach will do calculate each multi
func (m *Matrix) MultiEach(num interface{}) (matrix *Matrix) {
	var err error
	matrix = Copy(m)
	if m.CalcErr() != nil {
		return
	}
	if mat, ok := num.(Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.multiEachByMatrix(mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "MultiEach"}
		}
		return
	} else if mat, ok := num.(*Matrix); ok {
		if mat.CalcErr() != nil {
			matrix.calcErr = mat.calcErr
			return
		}
		if err = matrix.multiEachByMatrix(*mat); err != nil {
			matrix.calcErr = &MatrixError{msg: err.Error(), funcName: "MultiEach"}
		}
		return
	} else if mat, ok := num.(float64); ok {
		matrix.multiByFloat(float64(mat))
		return
	} else if mat, ok := num.(int); ok {
		matrix.multiByFloat(float64(mat))
		return
	}
	matrix.calcErr = &MatrixError{msg: "The multi op2 type is not allowed", funcName: "MultiEach"}
	return
}

// Sep
