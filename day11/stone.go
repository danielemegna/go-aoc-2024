package day11

import (
	"strconv"
)

type Stone int

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
	var left, _ = strconv.Atoi(engravedNumberString[:halfNumberOfDigits])
	var right, _ = strconv.Atoi(engravedNumberString[halfNumberOfDigits:])
	var rightStone = Stone(right)
	return Stone(left), &rightStone
}
