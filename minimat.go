package minimat

import (
	"errors"
)

type Mat struct {
	value [][]float64
	row   int
	col   int
}

func NewMat(row, col int) (m *Mat, err error) {
	if row < 1 || col < 1 {
		err = errors.New("error")
		return
	}
	v := make([][]float64, row)
	for i := 0; i < row; i++ {
		v[i] = make([]float64, col)
	}
	m = &Mat{value: v, row: row, col: col}
	return
}
func (m *Mat) Row() int { return m.row }
func (m *Mat) Col() int { return m.col }
func (m *Mat) Get(i, j int) (float64, error) {
	if !(0 <= i && i < m.row && 0 <= j && j < m.col) {
		return float64(0), errors.New("error")
	}
	return m.value[i][j], nil
}
func (m *Mat) Set(i, j int, v float64) error {
	if !(0 <= i && i < m.row && 0 <= j && j < m.col) {
		return errors.New("error")
	}
	m.value[i][j] = v
	return nil
}
func (m *Mat) Clone() *Mat {
	c, _ := NewMat(m.row, m.col)
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			c.value[i][j] = m.value[i][j]
		}
	}
	return c
}
func (a *Mat) Equal(b *Mat) bool {
	if b == nil || a.row != b.row || a.col != b.col {
		return false
	}
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			if a.value[i][j] != b.value[i][j] {
				return false
			}
		}
	}
	return true
}
func (a *Mat) Add(b *Mat) (*Mat, error) {
	if b == nil || a.row != b.row && a.col != b.col {
		return nil, errors.New("error")
	}
	c, _ := NewMat(a.row, a.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			c.value[i][j] = a.value[i][j] + b.value[i][j]
		}
	}
	return c, nil
}
func (a *Mat) Sub(b *Mat) (*Mat, error) {
	if b == nil || a.row != b.row && a.col != b.col {
		return nil, errors.New("error")
	}
	c, _ := NewMat(a.row, a.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			c.value[i][j] = a.value[i][j] - b.value[i][j]
		}
	}
	return c, nil
}
func (a *Mat) Dot(b *Mat) (*Mat, error) {
	if b == nil || a.row != b.col {
		return nil, errors.New("error")
	}
	c, _ := NewMat(a.row, b.col)
	for i := 0; i < a.row; i++ {
		for k := 0; k < a.col; k++ {
			for j := 0; j < b.col; j++ {
				c.value[i][j] += a.value[i][k] * b.value[k][j]
			}
		}
	}
	return c, nil
}
func (m *Mat) Pow(n int) (*Mat, error) {
	if m.row != m.col || n < 1 {
		return nil, errors.New("error")
	}
	c, _ := NewMat(m.row, m.col)
	for i := 0; i < m.row; i++ {
		c.value[i][i] = float64(1)
	}
	base := m.Clone()
	for i := 0; i < n; i++ {
		c, _ = c.Dot(base)
	}
	return c, nil
}
