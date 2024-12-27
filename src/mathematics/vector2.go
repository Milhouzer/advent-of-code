package mathematics

import (
	"math"
	"math/big"
)

type Vector2 struct {
	X float64
	Y float64
}

// Add other to a [Vector2]
func (v *Vector2) Add(other *Vector2) *Vector2 {
	return &Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// Remove other to a [Vector2]
func (v *Vector2) Remove(other *Vector2) *Vector2 {
	return &Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v *Vector2) Manhattan() int {
	return int(math.Abs(float64(v.X))) + int(math.Abs(float64(v.Y)))
}

func (v *Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2) InBounds(i, j int) bool {
	return v.X >= 0 && v.X < float64(i) && v.Y >= 0 && v.Y < float64(j)
}

// Vector2BigFloat uses big.Float for arbitrary precision
type Vector2BigFloat struct {
	X *big.Float
	Y *big.Float
}

// NewVector2BigFloat creates a new Vector2BigFloat from two float64 values
func NewVector2BigFloat(i, j float64) *Vector2BigFloat {
	return &Vector2BigFloat{
		X: big.NewFloat(i),
		Y: big.NewFloat(j),
	}
}

// Add adds another Vector2BigFloat to the current one
func (v *Vector2BigFloat) Add(other *Vector2BigFloat) *Vector2BigFloat {
	return &Vector2BigFloat{
		X: new(big.Float).Add(v.X, other.X),
		Y: new(big.Float).Add(v.Y, other.Y),
	}
}

// Remove subtracts another Vector2BigFloat from the current one
func (v *Vector2BigFloat) Remove(other *Vector2BigFloat) *Vector2BigFloat {
	return &Vector2BigFloat{
		X: new(big.Float).Sub(v.X, other.X),
		Y: new(big.Float).Sub(v.Y, other.Y),
	}
}

// Manhattan calculates the Manhattan distance (|I| + |J|) for the vector.
func (v *Vector2BigFloat) Manhattan() *big.Int {
	absI := new(big.Float).Abs(v.X)
	absJ := new(big.Float).Abs(v.Y)

	sum := new(big.Float).Add(absI, absJ)
	result, _ := sum.Int(nil)
	return result
}

// IsZero checks if the vector is the zero vector (both components are zero).
func (v *Vector2BigFloat) IsZero() bool {
	return v.X.Cmp(big.NewFloat(0)) == 0 && v.Y.Cmp(big.NewFloat(0)) == 0
}

// InBounds checks if the vector is within the bounds of i and j.
func (v *Vector2BigFloat) InBounds(i, j int) bool {
	bigI := big.NewFloat(float64(i))
	bigJ := big.NewFloat(float64(j))

	return v.X.Cmp(big.NewFloat(0)) >= 0 && v.X.Cmp(bigI) < 0 &&
		v.Y.Cmp(big.NewFloat(0)) >= 0 && v.Y.Cmp(bigJ) < 0
}
