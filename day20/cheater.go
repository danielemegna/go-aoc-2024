package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap, maxCheatDuration int) []Cheat {
	var result = []Cheat{}

	var cheatStartingElement = racetrackMap.RacetrackStart()
	for cheatStartingElement.Next != nil {
		var cheatStartDistanceFromStart = cheatStartingElement.DistanceFromStart

		for cheatEndingElement := cheatStartingElement.Next; cheatEndingElement != nil; cheatEndingElement = cheatEndingElement.Next {
			var cheatEndDistanceFromStart = cheatEndingElement.DistanceFromStart
			var cheatDuration = CalcManhattanDistance(cheatStartingElement.Coordinate, cheatEndingElement.Coordinate)
			if cheatDuration > maxCheatDuration {
				continue
			}

			var savedPicoseconds = cheatEndDistanceFromStart - cheatStartDistanceFromStart - cheatDuration
			if savedPicoseconds < 1 {
				continue
			}

			result = append(result, Cheat{
				startPosition:       cheatStartingElement.Coordinate,
				endPosition:         cheatEndingElement.Coordinate,
				savingInPicoseconds: savedPicoseconds,
			})
		}

		cheatStartingElement = cheatStartingElement.Next
	}

	return result
}
