package day19

type TowelPattern string

func (this TowelPattern) IsDesignPossibleWith(availableTowelPatterns AvailableTowelPatterns) bool {
	if availableTowelPatterns.IsAvailable(this) {
		return true
	}

	var firstColor = TowelPattern(this[0])
	if !availableTowelPatterns.IsAvailable(firstColor) {
		return false
	}

	var rest TowelPattern = this[1:]
	return rest.IsDesignPossibleWith(availableTowelPatterns)
}
