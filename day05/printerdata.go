package day05

import "github.com/samber/lo"

type PageOrderingRule struct {
	before int
	after  int
}

type PageOrderingRules []PageOrderingRule

func (this PageOrderingRules) RulesFor(pageNumber int) PageOrderingRules {
	return lo.Filter(this, func(rule PageOrderingRule, _ int) bool {
		return rule.before == pageNumber || rule.after == pageNumber
	})
}

type PagesToProduceInTheUpdate []int
