package day19

import (
	"regexp"
	"slices"
	"strings"
)

type TowelPattern = string
type ColorByte = byte

type AvailableTowelPatterns struct {
	isDesignPossibleRegex               *regexp.Regexp
	patternsByFirstColorDescSorted      map[ColorByte][]TowelPattern
	possibleDesignCombinationsByPattern map[TowelPattern]int
}

func AvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var stringRegexp = "^(" + strings.ReplaceAll(raw, ", ", "|") + ")+$"
	var regex = regexp.MustCompile(stringRegexp)

	var availableTowelPatterns = AvailableTowelPatterns{
		isDesignPossibleRegex:               regex,
		patternsByFirstColorDescSorted:      make(map[ColorByte][]TowelPattern),
		possibleDesignCombinationsByPattern: make(map[TowelPattern]int),
	}

	for rawPattern := range strings.SplitSeq(raw, ", ") {
		availableTowelPatterns.addDescSortedByLength(rawPattern)
	}

	return availableTowelPatterns
}

func (this AvailableTowelPatterns) IsAvailable(towelPattern TowelPattern) bool {
	var firstColorByte ColorByte = towelPattern[0]
	var existingSlice, keyExists = this.patternsByFirstColorDescSorted[firstColorByte]
	if !keyExists {
		return false
	}
	return slices.Contains(existingSlice, towelPattern)
}

func (this AvailableTowelPatterns) MaxPatternLengthFor(colorByte ColorByte) int {
	var existingSlice, keyExists = this.patternsByFirstColorDescSorted[colorByte]
	if !keyExists || len(existingSlice) == 0 {
		return 0
	}
	return len(existingSlice[0])
}

func (this AvailableTowelPatterns) IsDesignPossible(towelPattern TowelPattern) bool {
	var matches = this.isDesignPossibleRegex.FindStringSubmatch(towelPattern)
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

	var partialPatternLength = this.MaxPatternLengthFor(towelPattern[0])
	if partialPatternLength > len(towelPattern) {
		partialPatternLength = len(towelPattern)
	}

	var possibleCombination = 0
	for ; partialPatternLength > 0; partialPatternLength-- {
		var partialPattern = towelPattern[:partialPatternLength]
		var rest = towelPattern[partialPatternLength:]

		if !this.IsAvailable(partialPattern) {
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

func (this *AvailableTowelPatterns) addDescSortedByLength(rawStringPattern string) {
	var firstColorByte = rawStringPattern[0]
	var existingSlice, keyExists = this.patternsByFirstColorDescSorted[firstColorByte]
	if !keyExists {
		this.patternsByFirstColorDescSorted[firstColorByte] = []string{rawStringPattern}
		return
	}

	for index, existingElement := range existingSlice {
		if len(rawStringPattern) >= len(existingElement) {
			this.patternsByFirstColorDescSorted[firstColorByte] = slices.Insert(existingSlice, index, rawStringPattern)
			return
		}
	}
	this.patternsByFirstColorDescSorted[firstColorByte] = append(existingSlice, rawStringPattern)
}
