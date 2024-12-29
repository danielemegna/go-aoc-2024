package day05

func ParsePrinterData(inputLines []string) ([]PageOrderingRule, []PagesToProduceInTheUpdate) {
	var pageOrderingRule = []PageOrderingRule{
		{before: 47, after: 53},
		{before: 53, after: 13},
	}
	var pagesToProduceInTheUpdate = []PagesToProduceInTheUpdate{
		{75, 47, 61, 53, 29},
		{61, 13, 29},
	}
	return pageOrderingRule, pagesToProduceInTheUpdate
}
