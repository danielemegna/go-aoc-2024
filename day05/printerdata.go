package day05

type PageOrderingRule struct {
	before int
	after  int
}

type PageOrderingRules []PageOrderingRule

type PagesToProduceInTheUpdate []int
