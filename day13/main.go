package day13

func TotalTokensNeededToWinAllThePrizes(fileContent string) int {
	var machines = parseClawMachines(fileContent)
	var tokensNeeded = 0

	for _, machine := range machines {
		var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()
		if buttonAPressCount == -1 || buttonBPressCount == -1 {
			continue
		}

		tokensNeeded += 3 * buttonAPressCount
		tokensNeeded += buttonBPressCount
	}

	return tokensNeeded
}

func RealTotalTokensNeededToWinAllThePrizes(fileContent string) int {
	var machines = parseClawMachines(fileContent)
	var tokensNeeded = 0

	for _, machine := range machines {
		// TODO change the machine prize coordinates here ...
		var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()
		if buttonAPressCount == -1 || buttonBPressCount == -1 {
			continue
		}

		tokensNeeded += 3 * buttonAPressCount
		tokensNeeded += buttonBPressCount
	}

	return tokensNeeded
}
