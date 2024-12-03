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
	var fileBytes, _ = os.ReadFile("input.txt")
	var actual = TotalDistanceBetweenListsElements(string(fileBytes))
	assert.Equal(t, 1530215, actual)
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	t.Skip("WIP")
	var input =
		"3   4\n" +
		"4   3\n" +
		"2   5\n" +
		"1   3\n" +
		"3   9\n" +
		"3   3\n"
	assert.Equal(t, 31, SimilarityScoreFor(input))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileBytes, _ = os.ReadFile("input.txt")
	var actual = SimilarityScoreFor(string(fileBytes))
	assert.Equal(t, 9999, actual)
}
