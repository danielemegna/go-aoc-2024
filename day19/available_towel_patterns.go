package day19

import (
	"slices"
	"strings"
)

type AvailableTowelPatterns map[byte][]string

func InitAvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var patterns = AvailableTowelPatterns{}
	for pattern := range strings.SplitSeq(raw, ", ") {
		patterns.add(pattern)
	}

	return patterns
}

func (this AvailableTowelPatterns) IsAvailable(pattern string) bool {
	var firstColorByte = pattern[0]
	var patterns, keyExists = this[firstColorByte]
	return keyExists && slices.Contains(patterns, pattern)
}

func (this AvailableTowelPatterns) add(pattern string) {
	var firstColorByte = pattern[0]
	var existingSlice, keyExists = this[firstColorByte]
	if !keyExists {
		this[firstColorByte] = []string{pattern}
		return
	}

	for index, existingElement := range existingSlice {
		if len(pattern) >= len(existingElement) {
			this[firstColorByte] = slices.Insert(existingSlice, index, pattern)
			return
		}
	}
	this[firstColorByte] = append(existingSlice, pattern)
}

func (this AvailableTowelPatterns) MaxPatternLengthFor(pattern string) int {
	var firstColorByte = pattern[0]
	var patterns, keyExists = this[firstColorByte]
	if !keyExists {
		return 0
	}

	return len(patterns[0])
}
