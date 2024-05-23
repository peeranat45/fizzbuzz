package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBu(t *testing.T) {
	t.Run("input 1 expected 1", func(t *testing.T) {
		// Arrange
		expected := "1"
		input := "1"
		// Act
	
		output := fizzBuzz(input)

		// Assert

		assert.Equal(t, expected, output)
	})

	t.Run("input 2 expected 2", func(t *testing.T) {
		// Arrange 
		expected := "2"
		input := "2"
		// Act

		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, expected, output)
	})

	t.Run("input 3 expected Fizz", func(t *testing.T) {
		// Arrange 
		expected := "Fizz"
		input := "3"
		// Act

		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, expected, output)
	})

	t.Run("input 4 expected 4", func(t *testing.T) {
		// Arrange
		expected := "4"
		input := "4"

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})
}