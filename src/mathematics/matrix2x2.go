package mathematics

import (
	"errors"
	"math/big"
)

// [ X1 X2 ]
// [ Y1 Y2 ]
type Matrix2x2 struct {
	X1 float64
	X2 float64
	Y1 float64
	Y2 float64
}

var ErrNotInvertible = errors.New("matrix is not invertible")

func (m *Matrix2x2) Invert() (Matrix2x2, error) {
	det := m.Det()
	if det == 0 {
		return Matrix2x2{}, ErrNotInvertible
	}

	mat := Matrix2x2{
		X1: m.Y2,
		X2: -m.X2,
		Y1: -m.Y1,
		Y2: m.X1,
	}
	return mat.MulFloat(1 / det), nil
}

func (m *Matrix2x2) Det() float64 {
	return m.X1*m.Y2 - m.X2*m.Y1
}

func (m *Matrix2x2) MulFloat(s float64) Matrix2x2 {
	return Matrix2x2{
		X1: m.X1 * s,
		X2: m.X2 * s,
		Y1: m.Y1 * s,
		Y2: m.Y2 * s,
	}
}

func (m *Matrix2x2) MulVector2(v Vector2) Vector2 {
	return Vector2{
		X: m.X1*v.X + m.X2*v.Y,
		Y: m.Y1*v.X + m.Y2*v.Y,
	}
}

// [ X1 X2 ]
// [ Y1 Y2 ]
type Matrix2x2BigFloat struct {
	X1 *big.Float
	X2 *big.Float
	Y1 *big.Float
	Y2 *big.Float
}

func BigFromMatrix2x2(other Matrix2x2) Matrix2x2BigFloat {
	return Matrix2x2BigFloat{
		X1: big.NewFloat(other.X1).SetPrec(128),
		X2: big.NewFloat(other.X2).SetPrec(128),
		Y1: big.NewFloat(other.Y1).SetPrec(128),
		Y2: big.NewFloat(other.Y2).SetPrec(128),
	}
}

// Invert the matrix, using big.Float for arbitrary precision.
func (m *Matrix2x2BigFloat) Invert() (Matrix2x2BigFloat, error) {
	det := m.Det()
	if det.Cmp(big.NewFloat(0)) == 0 {
		return Matrix2x2BigFloat{}, ErrNotInvertible
	}

	invMat := Matrix2x2BigFloat{
		X1: m.Y2,
		X2: new(big.Float).Neg(m.X2),
		Y1: new(big.Float).Neg(m.Y1),
		Y2: m.X1,
	}

	detInv := new(big.Float).Quo(big.NewFloat(1), det)
	return invMat.MulFloat(detInv), nil
}

// Calculate the determinant of the matrix.
func (m *Matrix2x2BigFloat) Det() *big.Float {
	det := new(big.Float).Sub(
		new(big.Float).Mul(m.X1, m.Y2),
		new(big.Float).Mul(m.X2, m.Y1),
	)
	return det
}

// Multiply the matrix by a big.Float scalar.
func (m *Matrix2x2BigFloat) MulFloat(s *big.Float) Matrix2x2BigFloat {
	return Matrix2x2BigFloat{
		X1: new(big.Float).Mul(m.X1, s),
		X2: new(big.Float).Mul(m.X2, s),
		Y1: new(big.Float).Mul(m.Y1, s),
		Y2: new(big.Float).Mul(m.Y2, s),
	}
}

// Multiply the matrix by a Vector2.
func (m *Matrix2x2BigFloat) MulVector2(v Vector2BigFloat) Vector2BigFloat {
	return Vector2BigFloat{
		X: new(big.Float).Add(
			new(big.Float).Mul(m.X1, v.X),
			new(big.Float).Mul(m.X2, v.Y),
		),
		Y: new(big.Float).Add(
			new(big.Float).Mul(m.Y1, v.X),
			new(big.Float).Mul(m.Y2, v.Y),
		),
	}
}
