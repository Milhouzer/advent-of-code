package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStones(t *testing.T) {
	d := &Day11{}

	s1, s2, useSecond := d.getStones(125)
	assert.Equal(t, 253000, s1)
	assert.Equal(t, 0, s2)
	assert.False(t, useSecond)

	s1, s2, useSecond = d.getStones(512072)
	assert.Equal(t, 512, s1)
	assert.Equal(t, 72, s2)
	assert.True(t, useSecond)

	s1, s2, useSecond = d.getStones(512000)
	assert.Equal(t, 512, s1)
	assert.Equal(t, 0, s2)
	assert.True(t, useSecond)

	s1, s2, useSecond = d.getStones(6543127)
	assert.Equal(t, 6543127*2024, s1)
	assert.Equal(t, 0, s2)
	assert.False(t, useSecond)

	s1, s2, useSecond = d.getStones(20)
	assert.Equal(t, 2, s1)
	assert.Equal(t, 0, s2)
	assert.True(t, useSecond)

	s1, s2, useSecond = d.getStones(17)
	assert.Equal(t, 1, s1)
	assert.Equal(t, 7, s2)
	assert.True(t, useSecond)

	s1, s2, useSecond = d.getStones(1)
	assert.Equal(t, 2024, s1)
	assert.Equal(t, 0, s2)
	assert.False(t, useSecond)
}
