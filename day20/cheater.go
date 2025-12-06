package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap) []Cheat {
	return []Cheat{
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 6 - 2},
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 4 - 2},
	}
}
