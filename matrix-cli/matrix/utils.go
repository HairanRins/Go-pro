package matrix

import "errors"

// Add performs matrix addition
func Add(m1, m2 Matrix) (Matrix, error) {
	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
		return Matrix{}, errors.New("les dimensions des matrices doivent être identiques")
	}
	result := New(m1.Rows, m1.Cols, 0)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			result.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
		}
	}
	return result, nil
}

// Multiply performs matrix multiplication
func Multiply(m1, m2 Matrix) (Matrix, error) {
	if m1.Cols != m2.Rows {
		return Matrix{}, errors.New("le nombre de colonnes de m1 doit correspondre au nombre de lignes de m2")
	}
	result := New(m1.Rows, m2.Cols, 0)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m2.Cols; j++ {
			for k := 0; k < m1.Cols; k++ {
				result.Data[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}
	return result, nil
}
