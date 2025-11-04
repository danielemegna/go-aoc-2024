package day19

import (
	"slices"
	"strings"
)

type AvailableTowelPatterns map[ColorByte][]string
type ColorByte = byte

func InitAvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var patterns = AvailableTowelPatterns{}
	for pattern := range strings.SplitSeq(raw, ", ") {
		patterns.add(pattern)
	}

	return patterns
}

func (this AvailableTowelPatterns) IsAvailable(towelPattern TowelPattern) bool {
	var firstColorByte = towelPattern[0]
	var patterns, keyExists = this[firstColorByte]
	return keyExists && slices.Contains(patterns, string(towelPattern))
}

func (this AvailableTowelPatterns) MaxPatternLengthFor(colorByte ColorByte) int {
	var patterns, keyExists = this[colorByte]
	if !keyExists {
		return 0
	}

	return len(patterns[0])
}

func (this AvailableTowelPatterns) IsStringPatternAvailable(stringPattern string) bool {
	return this.IsAvailable(TowelPattern(stringPattern))
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
