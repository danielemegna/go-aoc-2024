package day19

import (
	"regexp"
	"slices"
	"strings"
)

type TowelPattern = string
type ColorByte = byte

type AvailableTowelPatterns struct {
	regex                               *regexp.Regexp
	patterns                            map[ColorByte][]TowelPattern
	possibleDesignCombinationsByPattern map[TowelPattern]int
}

func AvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var stringRegexp = "^(" + strings.ReplaceAll(raw, ", ", "|") + ")+$"
	var regex = regexp.MustCompile(stringRegexp)
	var patterns = AvailableTowelPatterns{
		regex:                               regex,
		patterns:                            make(map[ColorByte][]TowelPattern),
		possibleDesignCombinationsByPattern: make(map[TowelPattern]int),
	}
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
	var cachedResult, cachedResultExists = this.possibleDesignCombinationsByPattern[towelPattern]
	if cachedResultExists {
		return cachedResult
	}

	if !this.IsDesignPossible(towelPattern) {
		this.possibleDesignCombinationsByPattern[towelPattern] = 0
		return 0
	}

	var firstColorByte ColorByte = towelPattern[0]
	var partialPatternLength = len(this.patterns[firstColorByte][0])
	if partialPatternLength > len(towelPattern) {
		partialPatternLength = len(towelPattern)
	}

	var possibleCombination = 0
	for ; partialPatternLength > 0; partialPatternLength-- {
		var partialPattern = towelPattern[:partialPatternLength]
		var rest = towelPattern[partialPatternLength:]

		if !slices.Contains(this.patterns[firstColorByte], partialPattern) {
			continue
		}

		if len(rest) > 0 {
			possibleCombination += this.PossibleDesignCombinationsFor(rest)
			continue
		}

		possibleCombination++
	}

	this.possibleDesignCombinationsByPattern[towelPattern] = possibleCombination
	return possibleCombination
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
