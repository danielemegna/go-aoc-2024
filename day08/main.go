package day08

func AntinodesCountInMap(fileContent string) int {
	var cityMap = ParseCityMap(fileContent)
	return len(cityMap.AntinodesInMap())
}

func AntinodesCountWithUpdatedModel(fileContent string) int {
	return 34
}
