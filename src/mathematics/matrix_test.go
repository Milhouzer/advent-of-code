package mathematics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDet(t *testing.T) {
	m := Matrix2x2{
		X1: 26,
		X2: 67,
		Y1: 66,
		Y2: 21,
	}

	det := m.Det()
	assert.Equal(t, float64(-3876), det)
}

func TestDivisors(t *testing.T) {
	num1 := int64(10000000012176)
	num2 := int64(10000000012748)
	gcd := GCD(num1, num2)
	divisors := Divisors(gcd)
	assert.Equal(t, []int64([]int64{1, 4, 2}), divisors)

	num1 = int64(10000000008400)
	num2 = int64(10000000005400)
	gcd = GCD(num1, num2)
	divisors = Divisors(gcd)
	assert.Equal(t, 0, divisors)
}
