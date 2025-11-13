package day08

func SolveFirstPart(fileContent string) int {
	var am = NewAntennasMapFrom(fileContent)
	return am.CountAntinodes()
}
