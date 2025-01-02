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

	pagesLoop:
	for pageIndex := 0; pageIndex < len(this)-1; pageIndex++ {
		var currentPage, rest = clone[pageIndex], clone[pageIndex+1:]
		var rulesForCurrentPage = rules.RulesWithAfter(currentPage)

		if len(rulesForCurrentPage) == 0 {
			continue
		}

		for _, rule := range rulesForCurrentPage {
			if !slices.Contains(rest, rule.before) {
				continue
			}

			var indexOfPageToSwap = slices.Index(clone, rule.before)
			clone[pageIndex], clone[indexOfPageToSwap] = clone[indexOfPageToSwap], clone[pageIndex]
			pageIndex--
			continue pagesLoop
		}

	}

	return clone
}
