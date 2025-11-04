package day19

type TowelPattern string

func (this TowelPattern) IsDesignPossibleWith(availableTowelPatterns AvailableTowelPatterns) bool {
	if availableTowelPatterns.IsAvailable(this) {
		return true
	}

	var maxPatternLength = availableTowelPatterns.MaxPatternLengthFor(this[0])
	if maxPatternLength == 0 {
		return false
	}

	var partialPatternLength = maxPatternLength
	if(partialPatternLength > len(this)) {
		partialPatternLength = len(this)
	}

	for ; partialPatternLength > 0 ; partialPatternLength-- {
		var partialPattern = this[:partialPatternLength]
		if !availableTowelPatterns.IsAvailable(partialPattern) {
			continue
		}

		var rest = this[partialPatternLength:]
		var isRestDesignPossible = rest.IsDesignPossibleWith(availableTowelPatterns)
		if isRestDesignPossible {
			return true
		}
	}

	return false
}
