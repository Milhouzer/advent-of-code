package mathematics

import (
	"math"
	"math/big"
)

type Vector2 struct {
	I float64
	J float64
}

// Add other to a [Vector2]
func (v *Vector2) Add(other *Vector2) *Vector2 {
	return &Vector2{I: v.I + other.I, J: v.J + other.J}
}

// Remove other to a [Vector2]
func (v *Vector2) Remove(other *Vector2) *Vector2 {
	return &Vector2{I: v.I - other.I, J: v.J - other.J}
}

func (v *Vector2) Manhattan() int {
	return int(math.Abs(float64(v.I))) + int(math.Abs(float64(v.J)))
}

func (v *Vector2) IsZero() bool {
	return v.I == 0 && v.J == 0
}

func (v *Vector2) InBounds(i, j int) bool {
	return v.I >= 0 && v.I < float64(i) && v.J >= 0 && v.J < float64(j)
}

// Vector2BigFloat uses big.Float for arbitrary precision
type Vector2BigFloat struct {
	I *big.Float
	J *big.Float
}

// NewVector2BigFloat creates a new Vector2BigFloat from two float64 values
func NewVector2BigFloat(i, j float64) *Vector2BigFloat {
	return &Vector2BigFloat{
		I: big.NewFloat(i),
		J: big.NewFloat(j),
	}
}

// Add adds another Vector2BigFloat to the current one
func (v *Vector2BigFloat) Add(other *Vector2BigFloat) *Vector2BigFloat {
	return &Vector2BigFloat{
		I: new(big.Float).Add(v.I, other.I),
		J: new(big.Float).Add(v.J, other.J),
	}
}

// Remove subtracts another Vector2BigFloat from the current one
func (v *Vector2BigFloat) Remove(other *Vector2BigFloat) *Vector2BigFloat {
	return &Vector2BigFloat{
		I: new(big.Float).Sub(v.I, other.I),
		J: new(big.Float).Sub(v.J, other.J),
	}
}

// Manhattan calculates the Manhattan distance (|I| + |J|) for the vector.
func (v *Vector2BigFloat) Manhattan() *big.Int {
	absI := new(big.Float).Abs(v.I)
	absJ := new(big.Float).Abs(v.J)

	sum := new(big.Float).Add(absI, absJ)
	result, _ := sum.Int(nil)
	return result
}

// IsZero checks if the vector is the zero vector (both components are zero).
func (v *Vector2BigFloat) IsZero() bool {
	return v.I.Cmp(big.NewFloat(0)) == 0 && v.J.Cmp(big.NewFloat(0)) == 0
}

// InBounds checks if the vector is within the bounds of i and j.
func (v *Vector2BigFloat) InBounds(i, j int) bool {
	bigI := big.NewFloat(float64(i))
	bigJ := big.NewFloat(float64(j))

	return v.I.Cmp(big.NewFloat(0)) >= 0 && v.I.Cmp(bigI) < 0 &&
		v.J.Cmp(big.NewFloat(0)) >= 0 && v.J.Cmp(bigJ) < 0
}
