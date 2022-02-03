package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpToFive(t *testing.T) {
	out := UpToLength(5)
	assert.ElementsMatch(t, out,
		[]string{"1", "2", "FIZZ", "4", "BUZZ"})
}

func TestUpToFifteen(t *testing.T) {
	out := UpToLength(15)
	assert.ElementsMatch(t, out,
		[]string{"1", "2", "FIZZ", "4", "BUZZ",
			"FIZZ", "7", "8", "FIZZ", "BUZZ",
			"11", "FIZZ", "13", "14", "FIZZBUZZ"})
}
