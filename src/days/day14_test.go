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
		{I: 2, J: 4},
		{I: 4, J: 1},
		{I: 6, J: -2},
		{I: 8, J: -5},
		{I: 10, J: -8},
		{I: 12, J: -11},
	}
	// Expected results
	expected := []Vector2{
		{I: 2, J: 4},
		{I: 4, J: 1},
		{I: 6, J: 5},
		{I: 8, J: 2},
		{I: 10, J: 6},
		{I: 1, J: 3},
	}

	// Apply wrapping to each test case
	for i, test := range testCases {
		wrapped := wrap(test, xa, xb, ya, yb)
		assert.Equal(t, expected[i], wrapped)
	}
}
