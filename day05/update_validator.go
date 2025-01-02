package day05

import "slices"

func (this PagesToProduceInTheUpdate) IsValidFor(rules PageOrderingRules) bool {
	for i := 0; i < len(this)-1; i++ {
		var currentPageNumber, after = this[i], this[i+1:]

		var rulesForCurrentPage = rules.RulesWithAfter(currentPageNumber)
		if len(rulesForCurrentPage) == 0 {
			continue
		}

		for _, rule := range rulesForCurrentPage {
			if slices.Contains(after, rule.before) {
				return false
			}
		}

	}
	return true
}

func (this PagesToProduceInTheUpdate) FixWith(rules PageOrderingRules) PagesToProduceInTheUpdate {
	var clone = make(PagesToProduceInTheUpdate, len(this))
	copy(clone, this)
	for i := 0; i < len(this)-1; i++ {
		var currentPageNumber = this[i]
		var rulesForCurrentPage = rules.RulesWithAfter(currentPageNumber)
		if len(rulesForCurrentPage) == 0 {
			continue
		}

		if this[i+1] == rulesForCurrentPage[0].before {
			clone[i], clone[i+1] = clone[i+1], clone[i]
		}
	}

	return clone
}
