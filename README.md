# matrix
This repository give you basic matrix calculation and operation.

# Example
```golang
matrix := matrix.NewMatrix(4, 5, nil)
// it will make 4 * 5 matrix and all value initialized by 0


matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
// it will make 2 * 3 matrix 
// it looks like
// 1 2 3
// 4 5 6

matrix2 := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Add(matrix2)
// if you add matrix2 you will make matrix
// 2 4 6
// 8 10 12

matrix2 = matrix.NewMatrix(3, 1, []float64{1, 2, 3})
matrix.Multi(matrix2)
// if you Multi matrix your matrix will be
// 14
// 32
```
# operation
### Create Matrix
```golang

matrix := matrix.NewMatrix(4, 5, nil)
// it will make 4 * 5 matrix and all value initialized by 0

matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
// it will make 2 * 3 matrix 
// it looks like
// 1 2 3
// 4 5 6
```
### Create Vector
```golang
matrix := matrix.NewVector([]float64{1, 2, 3})
// if you use NewVecotr you will get 
// 1
// 2
// 3
```

### Copy your matrix
```golang 
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 := matrix.Copy(matrix)
// if you copy matrix then matrix2 got
// 1 2 3
// 4 5 6
```

### Add row
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.AddRow([]float64{7, 8, 9})
// if you want to add row then you can use AddRow
// if you use AddRow then matrix will be 
// 1 2 3
// 4 5 6
// 7 8 9

matrix = matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.NewMatrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
matrix.AddRow(matrix2)
// if you set matrix it will be
// 1 2 3
// 4 5 6
// 7 8 9
// 10 11 12
// 13 14 15

matrix = matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.AddRow(7)
// if you set int or float64, then you will get
// 1 2 3
// 4 5 6
// 7 7 7
```

### Add row at head
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.AddRowHEAD([]float64{7, 8, 9})
// if you want to add row then you can use AddRow
// if you use AddRow then matrix will be 
// 7 8 9
// 1 2 3
// 4 5 6

matrix = matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.NewMatrix{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
matrix.AddRowHEAD(matrix2)
// if you set matrix it will be
// 7 8 9
// 10 11 12
// 13 14 15
// 1 2 3
// 4 5 6

matrix = matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.AddRowHead(7)
// if you set int or float64, then you will get
// 7 7 7
// 1 2 3
// 4 5 6
```

### Get size
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
row, column := matrix.Size()
// then you wil get
// row = 2
// column = 3
```

### Get at value
```golang 
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
val := matrix.At(1, 1)
// then you will get
// val = 1
val = matrix.At(2, 2)
// then you will get
// val = 5
```

### Set at value
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Set(2, 1, float64(3.5))
// then you will get
// 1 2 3
// 3.5 5 6
```

### Substitution B to A
```golang
A := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
B := matrix.NewMatrix(2, 2, []float64{1, 2, 3, 4})
A.SetMatrix(B)
// then A will be
// 1 2
// 3 4
```

### Transpose matrix
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Transpose()
// then matrix will be
// 1 4
// 2 5
// 3 6
```

### SepRow matrix
```golang
matrix := &Matrix{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
matrix2 := matrix.SepRow(2, 4)
// then you will get matrix2
// 4, 5, 6
// 7, 8, 9
// 10, 11, 12
```

### SepColumn matrix
```golang
matrix = &Matrix{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
matrix2 := matrix.SepColumn(2, 4)
// then you will get matrix2
// 2, 3, 4
// 8, 9, 10
// 14, 15, 16
```

### Show matrix
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Show()
// it will do standard output like this
// 1 2 3
// 4 5 6
```

# calculation
### Add
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Add(2)
// then matrix will be
// 3 4 5
// 6 7 8

matrix2 := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Add(matrix2)
// then matrix will be
// 2 4 6
// 8 10 12
```
### Sub
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Sub(2)
// then matrix will be
// -1 0 1
// 2 3 4

matrix2 := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Sub(matrix2)
// then matrix will be
// 0 0 0
// 0 0 0
```
### Multi
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Multi(2)
// then matrix will be
// it do the same move to Multi Each
// 2 4 6
// 8 10 12

matrix2 := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.NewMatrix(3, 1, []float64{1, 2, 3})
matrix.Multi(matrix2)
// if you Multi matrix your matrix will be
// 14
// 32
```
### Multi Each
```golang
matrix := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.MultiEach(2)
// then matrix will be
// it do the same move to Multi Each
// 2 4 6
// 8 10 12

matrix2 := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.NewMatrix(3, 1, []float64{1, 2, 3, 4, 5, 6})
matrix.MultiEach(matrix2)
// if you Multi matrix your matrix will be
// 1 4 9
// 12 25 36
```