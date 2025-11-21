package day08

func AntinodesCountInMap(fileContent string) int {
	var cityMap = ParseCityMap(fileContent, false)
	return len(cityMap.AntinodesInMap())
}

func AntinodesCountWithUpdatedModel(fileContent string) int {
	var cityMap = ParseCityMap(fileContent, true)
	return len(cityMap.AntinodesInMap())
}
