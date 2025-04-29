package day11

import (
	"github.com/samber/lo"
	"strconv"
)

type Stone struct {
	engravedNumber int64
}

func (this Stone) OnBlink() (Stone, *Stone) {
	if this.engravedNumber == 0 {
		return Stone{engravedNumber: 1}, nil
	}

	var engravedNumberString = strconv.Itoa(int(this.engravedNumber))
	var numberOfDigits = len(engravedNumberString)
	if numberOfDigits%2 != 0 {
		return Stone{engravedNumber: this.engravedNumber * 2024}, nil
	}

	var halfNumberOfDigits = numberOfDigits / 2
	var left, _ = strconv.ParseInt(engravedNumberString[:halfNumberOfDigits], 10, 64)
	var right, _ = strconv.ParseInt(engravedNumberString[halfNumberOfDigits:], 10, 64)
	return Stone{engravedNumber: left}, &Stone{engravedNumber: right}
}

func StonesOnBlink(stones []Stone) []Stone {
	return lo.Map(stones, func(stone Stone, index int) Stone {
		var left, _ = stone.OnBlink()
		return left
	})
}
