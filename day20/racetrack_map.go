package day20

type RacetrackMap struct {
	values [][]MapValue
	start  RacetrackElement
	length int
}

type RacetrackElement struct {
	coordinate Coordinate
	next       *RacetrackElement
}

type MapValue = int

const START MapValue = 0
const WALL MapValue = -1

func ParseRacetrack(inputLines []string) RacetrackMap {
	var mapValues = make([][]MapValue, len(inputLines))
	var racetrackStart = RacetrackElement{}

	for y := 0; y < len(inputLines); y++ {
		mapValues[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				mapValues[y][x] = WALL
			case 'S':
				mapValues[y][x] = START
				racetrackStart.coordinate = Coordinate{x, y}
			default:
				mapValues[y][x] = 1
			}
		}
	}

	var trackLength = 0
	var currentRacetrackElement = &racetrackStart
	for {
		var c = currentRacetrackElement.coordinate
		var nextCoordinate Coordinate
		if mapValues[c.Y][c.X+1] == 1 {
			nextCoordinate = Coordinate{c.X + 1, c.Y}
		} else if mapValues[c.Y][c.X-1] == 1 {
			nextCoordinate = Coordinate{c.X - 1, c.Y}
		} else if mapValues[c.Y+1][c.X] == 1 {
			nextCoordinate = Coordinate{c.X, c.Y + 1}
		} else if mapValues[c.Y-1][c.X] == 1 {
			nextCoordinate = Coordinate{c.X, c.Y - 1}
		} else {
			break
		}

		mapValues[nextCoordinate.Y][nextCoordinate.X] = trackLength+1
		var newRacetrackElement = RacetrackElement{coordinate: nextCoordinate}
		currentRacetrackElement.next = &newRacetrackElement
		currentRacetrackElement = currentRacetrackElement.next
		trackLength++
	}

	return RacetrackMap{
		values: mapValues,
		start:  racetrackStart,
		length: trackLength,
	}
}

func (this RacetrackMap) StartCoordinate() Coordinate {
	return this.start.coordinate
}

func (this RacetrackMap) ValueAt(c Coordinate) MapValue {
	return this.values[c.Y][c.X]
}

func (this RacetrackMap) Length() int {
	return this.length
}
