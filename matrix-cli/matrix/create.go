package matrix

import "errors"

// CreateWith initializes a matrix with specific data 
func CreateWith(data [][]int) (Matrix, error)  {
	rows := len(data)
	if rows == 0 {
		return Matrix{}, errors.New("les données ne doivent pas être vides")
	}
	cols := len(data[0])
	for _, row := range data {
		if len(row) != cols {
			return Matrix{}, errors.New("toutes les lignes doivent avoir la même longueur")
		}
	}
	return Matrix{Rows: rows, Cols: cols, Data: data}, nil
}