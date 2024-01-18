package customerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CustomError(t *testing.T) {
	// Arrange
	err := CustomError{Message: "message"}

	// Act
	result := err.Error()

	// Assert
	assert.Equal(t, "message", result)
}

func Test_ArgumentError(t *testing.T) {
	// Arrange
	err := ArgumentError{Message: "message", Argument: "argument"}

	// Act
	result := err.Error()

	// Assert
	assert.Equal(t, "message: argument", result)
}
