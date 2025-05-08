package day11

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 55312, StonesCountAfterTwentyfiveBlinks("125 17"))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFirstFileLine()
	assert.Equal(t, 186996, StonesCountAfterTwentyfiveBlinks(fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 65601038650482, StonesCountAfterSeventyfiveBlinks("125 17"))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFirstFileLine()
	assert.Equal(t, 221683913164898, StonesCountAfterSeventyfiveBlinks(fileContent))
}

func readFirstFileLine() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	var fileContent = string(fileBytes)
	var lines = strings.Split(fileContent, "\n")
	return lines[0]
}
