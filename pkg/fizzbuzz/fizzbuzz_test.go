package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpToFive(t *testing.T) {
	out := UpToLength(5)
	assert.ElementsMatch(t, out,
		[]string{"1", "2", "Fizz", "4", "Buzz"})
}

func TestUpToFifteen(t *testing.T) {
	out := UpToLength(15)
	assert.ElementsMatch(t, out,
		[]string{"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz"})
}
