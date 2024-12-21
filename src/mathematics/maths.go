package mathematics

import (
	"math"
	"math/big"
)

// GCD calculates the greatest common divisor using the Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Divisors returns a slice containing all divisors of a given number
func Divisors(n int64) []int64 {
	var divisors []int64
	// Iterate from 1 to the square root of n to find divisors
	for i := int64(1); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i { // Avoid adding the square root twice for perfect squares
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

var (
	emin, _ = new(big.Float).SetString("0.0001")
	emax, _ = new(big.Float).SetString("0.9999")
)

func ToIntIfNear(v *big.Float) (bool, *big.Float) {
	if v.IsInt() {
		// fmt.Println("v is int: ", v.String())
		return true, v
	}

	intPart, _ := v.Int(nil)

	intPartFloat := new(big.Float).SetInt(intPart)

	result := new(big.Float).Sub(v, intPartFloat)
	if result.Cmp(big.NewFloat(0.5)) > 0 && result.Cmp(emax) > 0 {
		return true, intPartFloat
	}
	if result.Cmp(big.NewFloat(0.5)) < 0 && result.Cmp(emin) < 0 {
		return true, intPartFloat
	}
	// fmt.Println("intPartFloat:", intPartFloat.String())
	// fmt.Println("res:", result.String())

	return false, intPartFloat
}
