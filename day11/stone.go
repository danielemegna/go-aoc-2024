package day11

import (
	"github.com/samber/lo"
	"strconv"
)

type Stone int64

func (this Stone) OnBlink() (Stone, *Stone) {
	if this == 0 {
		return 1, nil
	}

	var engravedNumberString = strconv.Itoa(int(this))
	var numberOfDigits = len(engravedNumberString)
	if numberOfDigits%2 != 0 {
		return this * 2024, nil
	}

	var halfNumberOfDigits = numberOfDigits / 2
	var left, _ = strconv.ParseInt(engravedNumberString[:halfNumberOfDigits], 10, 64)
	var right, _ = strconv.ParseInt(engravedNumberString[halfNumberOfDigits:], 10, 64)
	var leftStone = Stone(left)
	var rightStone = Stone(right)
	return leftStone, &rightStone
}

func StonesOnBlink(stones []Stone) []Stone {
	return lo.FlatMap(stones, func(stone Stone, index int) []Stone {
		var left, right = stone.OnBlink()
		if right != nil {
			return []Stone{left, *right}
		}
		return []Stone{left}
	})
}
