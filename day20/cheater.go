package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap) []Cheat {
	var result = []Cheat{}

	var currentRacetrackPosition = racetrackMap.RacetrackStart()

	for currentRacetrackPosition != nil {

		var currentRacetrackValue = racetrackMap.ValueAt(currentRacetrackPosition.Coordinate)

		for _, firstStepCheat := range currentRacetrackPosition.Coordinate.CloseCoordinates() {
			if racetrackMap.ValueAt(firstStepCheat) != WALL {
				continue
			}

			for _, secondStepCheat := range firstStepCheat.CloseCoordinates() {
				if secondStepCheat.IsOutOfBounds(racetrackMap.MapSize()) {
					continue
				}
				if secondStepCheat == currentRacetrackPosition.Coordinate {
					continue
				}
				var secondStepValue = racetrackMap.ValueAt(secondStepCheat)
				if secondStepValue == WALL {
					continue
				}
				if secondStepValue < currentRacetrackValue {
					continue
				}

				var savedPicoseconds = secondStepValue - currentRacetrackValue - 2
				if(savedPicoseconds < 1) {
					continue
				}

				result = append(result, Cheat{
					startPosition:       currentRacetrackPosition.Coordinate,
					endPosition:         secondStepCheat,
					savingInPicoseconds: savedPicoseconds,
				})
			}

		}

		currentRacetrackPosition = currentRacetrackPosition.Next
	}

	return result
}
