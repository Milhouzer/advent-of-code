package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	xa, xb := 0, 11 // Horizontal range
	ya, yb := 0, 7  // Vertical range

	// Test cases
	testCases := []Vector2{
		{X: 2, Y: 4},
		{X: 4, Y: 1},
		{X: 6, Y: -2},
		{X: 8, Y: -5},
		{X: 10, Y: -8},
		{X: 12, Y: -11},
	}
	// Expected results
	expected := []Vector2{
		{X: 2, Y: 4},
		{X: 4, Y: 1},
		{X: 6, Y: 5},
		{X: 8, Y: 2},
		{X: 10, Y: 6},
		{X: 1, Y: 3},
	}

	// Apply wrapping to each test case
	for i, test := range testCases {
		wrapped := wrap(test, xa, xb, ya, yb)
		assert.Equal(t, expected[i], wrapped)
	}
}
