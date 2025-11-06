package day19

import (
	"regexp"
	"strings"
)

type TowelPattern = string
type ColorByte = byte

type AvailableTowelPatterns struct {
	regex    *regexp.Regexp
	patterns []string
}

func AvailableTowelPatternsFrom(raw string) AvailableTowelPatterns {
	var stringRegexp = "^(" + strings.ReplaceAll(raw, ", ", "|") + ")+$"
	var regexp = regexp.MustCompile(stringRegexp)
	var patterns = strings.Split(raw, ", ")
	return AvailableTowelPatterns{regexp, patterns}
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
