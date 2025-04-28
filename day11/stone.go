package day11

import "strconv"

type Stone struct {
	engravedNumber int64
}

func (this Stone) OnBlink() (Stone, *Stone) {
	if this.engravedNumber == 0 {
		return Stone{engravedNumber: 1}, nil
	}

	var engravedNumberString = strconv.Itoa(int(this.engravedNumber))
	var length = len(engravedNumberString)
	var lengthHalf = length / 2

	var left, _ = strconv.ParseInt(engravedNumberString[:lengthHalf], 10, 64)
	var right, _ = strconv.ParseInt(engravedNumberString[lengthHalf:], 10, 64)
	return Stone{engravedNumber: left}, &Stone{engravedNumber: right}
}
