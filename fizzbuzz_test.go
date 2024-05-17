package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBu(t *testing.T) {
	t.Run("input 1 expected 1", func(t *testing.T) {
		// Arrange
		expected := "1"

		// Act
	
		output := fizzBuzz(expected)

		// Assert

		assert.Equal(t, expected, output)
	})
}