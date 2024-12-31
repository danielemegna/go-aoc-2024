package day05

func (this PagesToProduceInTheUpdate) IsValidFor(rules PageOrderingRules) bool {
	for i := 1; i < len(this)-1; i++ {
		var before, currentPageNumber, after = this[:i], this[i], this[i+1:]

		var rulesForCurrentPage = rules.RulesFor(currentPageNumber)
		if len(rulesForCurrentPage) == 0 {
			return true
		}

		if rulesForCurrentPage[0].before == after[0] {
			return false
		}

		if rulesForCurrentPage[0].after == before[0] {
			return false
		}

	}
	return true
}
