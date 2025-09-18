package day17

import (
	"strconv"
)

func BitwiseXor(a int64, b int64) int {
	// TODO handle b > a

	var binaryA = toBinaryArray(a)
	var binaryB = toBinaryArray(b)

	var result = []bool{}
	for i := range binaryA {
		var bitB = false
		if len(binaryB) > i {
			bitB = binaryB[i]
		}
		var bit = binaryA[i] != bitB
		result = append(result, bit)
	}

	return toInt(result)
}

func toBinaryArray(n int64) []bool {
	var stringBinary = strconv.FormatInt(n, 2)

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

func toInt(binaryArray []bool) int {
	// TODO len(binaryArray) > 1
	if binaryArray[0] {
		return 1
	}

	return 0
}
