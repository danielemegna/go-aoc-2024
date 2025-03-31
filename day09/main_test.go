package day09

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 1928, FilesystemChecksumAfterDefrag("2333133121414131402"))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 6340197768906, FilesystemChecksumAfterDefrag(fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 2858, FilesystemChecksumAfterDefragWholeFiles("2333133121414131402"))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 6_363_913_128_533, FilesystemChecksumAfterDefragWholeFiles(fileContent))
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
