package day17

import (
	"math"
	"strconv"
)

type BinaryNumber []bool

func BinaryNumberFromInt(number int) BinaryNumber {
	var stringBinary = strconv.FormatInt(int64(number), 2)

	var result = []bool{}
	for _, char := range stringBinary {
		switch char {
		case '0':
			result = append(result, false)
		case '1':
			result = append(result, true)
		}
	}

	return result
}

func (this BinaryNumber) ToInt() int {
	var result = 0
	var bottomIndex = len(this) - 1
	for index := range this {
		if this[bottomIndex] {
			result += int(math.Pow(float64(2), float64(index)))
		}
		bottomIndex--
	}

	return result
}

func (this BinaryNumber) PadLeft(totalSize int) BinaryNumber {
	if len(this) >= totalSize {
		return this
	}

	var pad = make([]bool, totalSize-len(this))
	return append(pad, this...)
}
