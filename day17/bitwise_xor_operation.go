package day17

func BitwiseXor(a int, b int) int {
	// discovered @ SoCraTes Italy 2025 discussing about the problem with
	// @gabrielelana that go already provide a set of Arithmetic operators
	// https://go.dev/ref/spec#Arithmetic_operators
	// ... ManualBitwiseXor and Binary number become not needed !
	return a ^ b
}

func ManualBitwiseXor(a int, b int) int {
	var binaryA = BinaryNumberFromInt(a)
	var binaryB = BinaryNumberFromInt(b)
	var lengthA = len(binaryA)
	var lengthB = len(binaryB)

	if lengthA > lengthB {
		binaryB = binaryB.PadLeft(lengthA)
	} else if lengthB > lengthA {
		binaryA = binaryA.PadLeft(lengthB)
	}

	var result = BinaryNumber{}
	for i := range binaryA {
		var bit = binaryA[i] != binaryB[i]
		result = append(result, bit)
	}

	return result.ToInt()
}
