package day11

import (
	"strconv"
)

type Stone int

func (this Stone) OnBlink() (Stone, *Stone) {
	if this == 0 {
		return 1, nil
	}

	var numberOfDigits = len(this.toString())
	if numberOfDigits%2 != 0 {
		return this * 2024, nil
	}

	var leftStone = this.getLeftHalfStone()
	var rightStone = this.getRightHalfStone()
	return leftStone, &rightStone
}

func (this Stone) getLeftHalfStone() Stone {
	var engravedNumberString = this.toString()
	var halfNumberOfDigits = len(engravedNumberString) / 2
	var newStoneNumber, _ = strconv.Atoi(engravedNumberString[:halfNumberOfDigits])
	return Stone(newStoneNumber)
}

func (this Stone) getRightHalfStone() Stone {
	var engravedNumberString = this.toString()
	var halfNumberOfDigits = len(engravedNumberString) / 2
	var newStone, _ = strconv.Atoi(engravedNumberString[halfNumberOfDigits:])
	return Stone(newStone)
}

func (this Stone) toString() string {
	return strconv.Itoa(int(this))
}
