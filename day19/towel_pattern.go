package day19

type TowelPattern string

func (this TowelPattern) IsDesignPossibleWith(availableTowelPatterns AvailableTowelPatterns) bool {
	if availableTowelPatterns.IsAvailable(this) {
		return true
	}

	var maxPatternLength = availableTowelPatterns.MaxPatternLengthFor(this)
	if maxPatternLength == 0 {
		return false
	}

	var pieceSize = maxPatternLength
	if(maxPatternLength > len(this)) {
		pieceSize = len(this)
	}

	for ; pieceSize > 0 ; pieceSize-- {
		var patternPiece = this[:pieceSize]
		if !availableTowelPatterns.IsAvailable(patternPiece) {
			continue
		}

		var rest = this[pieceSize:]
		var isDesignPossible = rest.IsDesignPossibleWith(availableTowelPatterns)
		if isDesignPossible {
			return true
		}
	}

	return false
}
