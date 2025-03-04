package day09

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 1928, FilesystemChecksumAfterCompact("2333133121414131402"))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	assert.Equal(t, -1, FilesystemChecksumAfterCompact(fileContent))
}

func simulateFileContent(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
