package day01

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var input =
	  "3   4\n" +
		"4   3\n" +
		"2   5\n" +
		"1   3\n" +
		"3   9\n" +
		"3   3\n"
	assert.Equal(t, 11, TotalDistanceBetweenListsElements(input))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var bytes, _ = os.ReadFile("input.txt")
	var input = string(bytes)
	assert.Equal(t, 1530215, TotalDistanceBetweenListsElements(input))
}
