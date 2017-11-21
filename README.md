# matrix
This repository give you basic matrix calculation and operation.
If you change of oparate matrix oftenly they will not return error separetely.
When error happen function return matrix which has error in Matrix struct.
if you want to chekc error you can see by **matrix.Err()**

also if matrix.Err() != nil
matrix will not calculate.
thus if you finish bach of calculate please check matrix.Err()

# Example
```golang
matrix := matrix.New(4, 5, nil)
// it will make 4 * 5 matrix and all value initialized by 0


matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
// it will make 2 * 3 matrix 
// it looks like
// 1 2 3
// 4 5 6

matrix2 := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Add(matrix2)
// if you add matrix2 you will make matrix
// 2 4 6
// 8 10 12

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New(3, 1, []float64{1, 2, 3})
matrix = matrix.Multi(matrix2)
// if you Multi matrix your matrix will be
// 14
// 32

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New(1, 2, []float64{1, 2})
matrix = matrix.Multi(matrix2)
matrix.Err()
// it will show you that err happen
```
# operation
### Create Matrix
```golang

matrix := matrix.New(4, 5, nil)
// it will make 4 * 5 matrix and all value initialized by 0

matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
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

### Create Hot Vector
```golang
matrix := NewHotVector(5, 3)
// then you will get
// 0
// 0
// 1
// 0
// 0
```

### Create Random Matrix
```golang
matrix := matrix.NewRandom(2, 4, 3)
// the third argument mean number of decimal places.
// you will get(about)
// 0.123 0.234 0.345 0.124
// 0.646 0.683 0.573 0.198
```

## Create Eye Matrix
```golang
matrix := matrix.NewEye(3)
// then you will get
// 1 0 0
// 0 1 0
// 0 0 1
```

### Copy your matrix
```golang 
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 := matrix.Copy(matrix)
// if you copy matrix then matrix2 got
// 1 2 3
// 4 5 6
```

### Add row
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.AddRow([]float64{7, 8, 9})
// if you want to add row then you can use AddRow
// if you use AddRow then matrix will be 
// 1 2 3
// 4 5 6
// 7 8 9

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
matrix = matrix.AddRow(matrix2)
// if you set matrix it will be
// 1 2 3
// 4 5 6
// 7 8 9
// 10 11 12
// 13 14 15

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.AddRow(7)
// if you set int or float64, then you will get
// 1 2 3
// 4 5 6
// 7 7 7
```

### Add row at head
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.AddRowHEAD([]float64{7, 8, 9})
// if you want to add row then you can use AddRow
// if you use AddRow then matrix will be 
// 7 8 9
// 1 2 3
// 4 5 6

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New{3, 3, []float64{7, 8, 9, 10, 11, 12, 13, 14, 15}}
matrix = matrix.AddRowHEAD(matrix2)
// if you set matrix it will be
// 7 8 9
// 10 11 12
// 13 14 15
// 1 2 3
// 4 5 6

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.AddRowHead(7)
// if you set int or float64, then you will get
// 7 7 7
// 1 2 3
// 4 5 6
```

### Get size
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
row, column := matrix.Size()
// then you wil get
// row = 2
// column = 3

row = matrix.Row()
// then you will get
// 2

column = matrix.Column()
// then you will get 
// 3
```

### Get at value
```golang 
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
val, err := matrix.At(1, 1)
// then you will get
// val = 1
val, err = matrix.At(2, 2)
// then you will get
// val = 5
```

### Set at value
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
err = matrix.Set(2, 1, float64(3.5))
// then you will get
// 1 2 3
// 3.5 5 6
```

### Reshape
```golang
matrix := matrix.New(6, 1, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Reshape(2, 3)
// then you will get
// 1 2 3
// 4 5 6
```

### Substitution B to A
```golang
A := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
B := matrix.New(2, 2, []float64{1, 2, 3, 4})
A.SetMatrix(B)
// then A will be
// 1 2
// 3 4
```

### Transpose matrix
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Transpose()
// then matrix will be
// 1 4
// 2 5
// 3 6
```

### SepRow matrix
```golang
matrix := matrix.New{6, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
matrix = matrix.SepRow(2, 4)
// then you will get matrix
// 4, 5, 6
// 7, 8, 9
// 10, 11, 12
```

### SepColumn matrix
```golang
matrix := matrix.New{3, 6, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}}
matrix = matrix.SepColumn(2, 4)
// then you will get matrix
// 2, 3, 4
// 8, 9, 10
// 14, 15, 16
```

### Vector
```golang
matrix := matrix.New{2, 2, []float64{1, 2, 3, 4}}
matrix = matrix.Vector()
// then you will get
// 1
// 2
// 3
// 4
```

### Show matrix
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix.Show()
// it will do standard output like this
// 1 2 3
// 4 5 6
```

# calculation

### Add
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Add(2)
// then matrix will be
// 3 4 5
// 6 7 8

matrix2 := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Add(matrix2)
// then matrix will be
// 2 4 6
// 8 10 12
```
### Sub
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Sub(2)
// then matrix will be
// -1 0 1
// 2 3 4

matrix2 := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Sub(matrix2)
// then matrix will be
// 0 0 0
// 0 0 0
```
### Multi
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Multi(2)
// then matrix will be
// it do the same move to Multi Each
// 2 4 6
// 8 10 12

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New(3, 1, []float64{1, 2, 3})
matrix = matrix.Multi(matrix2)
// if you Multi matrix your matrix will be
// 14
// 32
```
### Multi Each
```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.MultiEach(2)
// then matrix will be
// it do the same move to Multi Each
// 2 4 6
// 8 10 12

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New(3, 1, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.MultiEach(matrix2)
// if you Multi matrix your matrix will be
// 1 4 9
// 12 25 36
```

### Div 
if you divin by 0 you will get Inf.

```golang
matrix := matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Div(0.1)
// then matrix will be 
// 10 20 30
// 40 50 60

matrix = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix2 = matrix.New(2, 3, []float64{1, 2, 3, 4, 5, 6})
matrix = matrix.Div(matrix2)
// then matrix will be 
// 1 1 1
// 1 1 1
```

### Sigmoid
```golang
matrix := matrix.New{3, 4, []float64{1, 2, 3, 4, 5, 6, -1, -2, -3}}
matrix = matrix.Sigmoid()
// then you will get (almost)
// 0.731059   0.880797   0.952574
// 0.982014   0.993307   0.997527
// 0.268941   0.119203   0.047426
```

### Sigmoid Gradient
```golang
matrix := New{3, 3, []float64{1, 2, 3, 4, 5, 6, -1, -2, -3}, nil}
matrix = matrix.SigmoidGradient()
// then you will get (almost)
// 0.1966119   0.1049936   0.0451767
// 0.0176627   0.0066481   0.0024665
// 0.1966119   0.1049936   0.0451767
```