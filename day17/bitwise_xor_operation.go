package day17

func BitwiseXor(a int, b int) int {
	// TODO handle b > a

	var binaryA = toBinaryArray(a)
	var binaryB = toBinaryArray(b)

	var result = []bool{}
	for i := range binaryA {
		var bit = binaryA[i] != binaryB[i]
		result = append(result, bit)
	}

	return toInt(result)
}

func toBinaryArray(n int) []bool {
	// TODO handle n > 1
	if n == 0 {
		return []bool{false}
	}

	return []bool{true}
}

func toInt(binaryArray []bool) int {
	// TODO len(binaryArray) > 1
	if binaryArray[0] {
		return 1
	}

	return 0
}
