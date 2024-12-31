package day05

import (
	"fmt"
	"github.com/samber/lo"
)

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

func (this PagesToProduceInTheUpdate) GetMiddlePageNumber() int {
	var size = len(this)
	if size%2 == 0 {
		panic(fmt.Sprintf("Cannot get middle page number for %d pages!", size))
	}

	return this[len(this)/2]
}
