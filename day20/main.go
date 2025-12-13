package day20

import "github.com/samber/lo"

func CountCheatsSavingAtLeast(picoseconds int, fileContent string) int {
	var racetrackMap = ParseRacetrack(fileContent)
	var possibleCheats = PossibleCheatsIn(racetrackMap, 2)
	return lo.CountBy(possibleCheats, func(cheat Cheat) bool {
		return cheat.savingInPicoseconds >= picoseconds
	})
}

func CountLongCheatsSavingAtLeast(picoseconds int, fileContent string) int {
	var racetrackMap = ParseRacetrack(fileContent)
	var possibleCheats = PossibleCheatsIn(racetrackMap, 20)
	return lo.CountBy(possibleCheats, func(cheat Cheat) bool {
		return cheat.savingInPicoseconds >= picoseconds
	})
}
