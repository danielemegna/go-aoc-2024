package day02

type Report struct {
	levels []int
}

func (this Report) IsValid() bool {
	var isAValidIncreasingReport = true
	var isAValidDecreasingReport = true

	for i := 0; i < len(this.levels)-1; i++ {
		var levelDifferenceWithNext = this.levels[i+1] - this.levels[i]

		if isAValidIncreasingReport {
			isAValidIncreasingReport = isAValidIncreasingDiffence(levelDifferenceWithNext)
		}

		if isAValidDecreasingReport {
			isAValidDecreasingReport = isAValidDecreasingDiffence(levelDifferenceWithNext)
		}

		if !isAValidIncreasingReport && !isAValidDecreasingReport {
			return false
		}
	}

	return true
}

func isAValidIncreasingDiffence(levelsDifference int) bool {
	return levelsDifference > 0 && levelsDifference <= 3
}

func isAValidDecreasingDiffence(levelsDifference int) bool {
	return levelsDifference < 0 && levelsDifference >= -3
}
