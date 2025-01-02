package day05

import "slices"

func (this PagesToProduceInTheUpdate) IsValidFor(rules PageOrderingRules) bool {
	for i := 0; i < len(this)-1; i++ {
		var currentPageNumber, rest = this[i], this[i+1:]
		var rulesForCurrentPage = rules.RulesWithAfter(currentPageNumber)

		if len(rulesForCurrentPage) == 0 {
			continue
		}

		for _, rule := range rulesForCurrentPage {
			if slices.Contains(rest, rule.before) {
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
		var currentPageNumber, rest = clone[i], clone[i+1:]
		var rulesForCurrentPage = rules.RulesWithAfter(currentPageNumber)

		if len(rulesForCurrentPage) == 0 {
			continue
		}

		for _, rule := range rulesForCurrentPage {
			var indexOf = slices.Index(rest, rule.before)
			if indexOf == -1 {
				continue
			}

			indexOf += i + 1
			clone[i], clone[indexOf] = clone[indexOf], clone[i]
		}

	}

	return clone
}
