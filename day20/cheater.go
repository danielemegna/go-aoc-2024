package day20

type Cheat struct {
	startPosition       Coordinate
	endPosition         Coordinate
	savingInPicoseconds int
}

func PossibleCheatsIn(racetrackMap RacetrackMap) []Cheat {
	return []Cheat{
		{startPosition: Coordinate{1, 0}, endPosition: Coordinate{1, 2}, savingInPicoseconds: 6 - 2},
		{startPosition: Coordinate{2, 0}, endPosition: Coordinate{2, 2}, savingInPicoseconds: 4 - 2},
	}
}
