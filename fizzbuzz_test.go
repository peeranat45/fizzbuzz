package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("input 1 expected 1", func(t *testing.T) {
		// Arrange
		expected := "1"
		input := 1
		// Act
	
		output := fizzBuzz(input)

		// Assert

		assert.Equal(t, expected, output)
	})

	t.Run("input 2 expected 2", func(t *testing.T) {
		// Arrange 
		expected := "2"
		input := 2
		// Act

		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, expected, output)
	})

	t.Run("input 3 expected Fizz", func(t *testing.T) {
		// Arrange 
		expected := "Fizz"
		input := 3
		// Act

		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, expected, output)
	})

	t.Run("input 4 expected 4", func(t *testing.T) {
		// Arrange
		expected := "4"
		input := 4

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})

	t.Run("input 5 expected Buzz", func(t *testing.T) {
		// Arrange
		expected := "Buzz"
		input := 5

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})

	t.Run("input 6 expected Fizz", func(t *testing.T) {
		// Arrange
		expected := "Fizz"
		input := 6

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})

	t.Run("input 9 expected Fizz", func(t *testing.T) {
		// Arrange
		expected := "Fizz"
		input := 9

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})

	t.Run("input 10 expected Buzz", func(t *testing.T) {
		// Arrange
		expected := "Buzz"
		input := 10

		// Act
		output := fizzBuzz(input)

		// Assert
		assert.Equal(t, output, expected)
	})
}