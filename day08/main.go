package day08

func AntinodesInMap(fileContent string) int {
	var cityMap = ParseCityMap(fileContent)
	return len(cityMap.AntinodesInMap())
}
