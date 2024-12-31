package day05

import "slices"

func (this PagesToProduceInTheUpdate) IsValidFor(rules PageOrderingRules) bool {
	for i := 0; i < len(this)-1; i++ {
		var before, currentPageNumber, after = this[:i], this[i], this[i+1:]

		var rulesForCurrentPage = rules.RulesFor(currentPageNumber)
		if len(rulesForCurrentPage) == 0 {
			continue
		}

		var rule = rulesForCurrentPage[0]

		if slices.Contains(after, rule.before) {
			return false
		}

		if len(before) == 0 {
			continue
		}

		if rule.after == before[0] {
			return false
		}

	}
	return true
}
