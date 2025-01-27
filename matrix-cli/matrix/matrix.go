package matrix 

import (
	"fmt"
)

// Matrix structure 
type Matrix struct {
	Rows, Cols int 
	Data		[][]int
}

// New creates a new matrix with default values
func New(rows, cols, defaultValue int) Matrix  {
	data := make([][]int, rows)
	for i := range data  {
		data[i] = make([]int, cols)
		for j := range data[i] {
			data[i][j] = defaultValue
		}
	}
	return Matrix{Rows: rows, Cols: cols, Data: data}
}

// Print displays the matrix 
func (m Matrix) Print()  {
	for _, row := range m.Data {
		fmt.Println(row)
	}
}

