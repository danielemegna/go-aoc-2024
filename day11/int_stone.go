package day11

import (
	"github.com/samber/lo"
	"strconv"
)

type iStone = int

func OnBlink(stone iStone) (iStone, *iStone) {
	if stone == 0 {
		return 1, nil
	}

	var engravedNumberString = strconv.Itoa(int(stone))
	var numberOfDigits = len(engravedNumberString)
	if numberOfDigits%2 != 0 {
		return stone * 2024, nil
	}

	var halfNumberOfDigits = numberOfDigits / 2
	var left, _ = strconv.Atoi(engravedNumberString[:halfNumberOfDigits])
	var right, _ = strconv.Atoi(engravedNumberString[halfNumberOfDigits:])
	return left, &right
}

func iStonesOnBlink(stones []iStone) []iStone {
	return lo.FlatMap(stones, func(stone iStone, index int) []iStone {
		var left, right = OnBlink(stone)
		if right != nil {
			return []iStone{left, *right}
		}
		return []iStone{left}
	})
}
