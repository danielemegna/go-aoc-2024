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
	var startCoordinate Coordinate

	for y := 0; y < len(inputLines); y++ {
		mapValues[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				mapValues[y][x] = WALL
			case 'S':
				mapValues[y][x] = START
				startCoordinate = Coordinate{x, y}
			default:
				mapValues[y][x] = 1
			}
		}
	}

	var trackLength = 8
	// TODO:
	// - complete linked list from start
	// - update mapValues with values > 1
	// - detect trackLength from map

	return RacetrackMap{
		values: mapValues,
		start:  RacetrackElement{coordinate: startCoordinate, next: nil},
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
