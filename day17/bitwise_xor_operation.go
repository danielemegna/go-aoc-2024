package day17

func BitwiseXor(a int64, b int64) int64 {
	var binaryA = BinaryNumberFromInt(a)
	var binaryB = BinaryNumberFromInt(b)

	// TODO handle b > a
	binaryB = binaryB.PadLeft(len(binaryA))

	var result = BinaryNumber{}

	for i := range binaryA {
		var bit = binaryA[i] != binaryB[i]
		result = append(result, bit)
	}

	return result.ToInt()
}
