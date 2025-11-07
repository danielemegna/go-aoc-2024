package day19

import (
	"regexp"
	"slices"
	"strings"
)

type TowelPattern = string
type ColorByte = byte

type AvailableTowelPatterns struct {
	regex    *regexp.Regexp
	patterns map[ColorByte][]TowelPattern
}

func AvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var stringRegexp = "^(" + strings.ReplaceAll(raw, ", ", "|") + ")+$"
	var regexp = regexp.MustCompile(stringRegexp)
	var patterns = AvailableTowelPatterns{regexp, map[ColorByte][]string{}}
	for pattern := range strings.SplitSeq(raw, ", ") {
		patterns.add(pattern)
	}

	return patterns
}

func (this AvailableTowelPatterns) IsDesignPossible(towelPattern TowelPattern) bool {
	var matches = this.regex.FindStringSubmatch(towelPattern)
	return len(matches) > 0 && matches[0] == towelPattern
}

func (this AvailableTowelPatterns) PossibleDesignCombinationsFor(towelPattern TowelPattern) int {
	if !this.IsDesignPossible(towelPattern) {
		return 0
	}

	return 1
}

func (this *AvailableTowelPatterns) add(rawStringPattern string) {
	var firstColorByte = rawStringPattern[0]
	var existingSlice, keyExists = this.patterns[firstColorByte]
	if !keyExists {
		this.patterns[firstColorByte] = []string{rawStringPattern}
		return
	}

	for index, existingElement := range existingSlice {
		if len(rawStringPattern) >= len(existingElement) {
			this.patterns[firstColorByte] = slices.Insert(existingSlice, index, rawStringPattern)
			return
		}
	}
	this.patterns[firstColorByte] = append(existingSlice, rawStringPattern)
}
