package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	// Arrange
	values := []int{1, 2, 3, 4, 5}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal(5, max)
	assert.Nil(err)
}

type Group int

func TestMaxGroup(t *testing.T) {
	// Arrange
	values := []Group{1, 2, 3, 4, 5}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal(Group(5), max)
	assert.Nil(err)
}

func TestMaxString(t *testing.T) {
	// Arrange
	values := []string{"a", "b", "c", "d", "e"}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal("e", max)
	assert.Nil(err)
}
func TestMaxIntError(t *testing.T) {
	// Arrange
	values := []int{}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal(0, max)
	assert.Error(err)
}

func TestMaxFloat(t *testing.T) {
	// Arrange
	values := []float64{1, 2, 3, 4, 5}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal(float64(5), max)
	assert.Nil(err)
}

func TestMaxError(t *testing.T) {
	// Arrange
	values := []float64{}

	// Act
	max, err := Max(values)

	// Assert
	assert := assert.New(t)
	assert.Equal(float64(0), max)
	assert.Error(err)
}
