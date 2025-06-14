package day12

import "github.com/samber/lo"

func TotalFencePrice(fileContent string) int {
	var gardenMap = ParseGardenMap(fileContent)

	return lo.SumBy(gardenMap, func(r GardenRegion) int {
		return r.area * r.perimeter.Length()
	})
}

func TotalFencePriceWithBulkDiscount(fileContent string) int {
	var gardenMap = ParseGardenMap(fileContent)

	return lo.SumBy(gardenMap, func(r GardenRegion) int {
		return r.area * r.perimeter.NumberOfSides()
	})
}
