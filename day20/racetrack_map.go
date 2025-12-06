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
	var endCoordinate Coordinate

	for y := 0; y < len(inputLines); y++ {
		mapValues[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				mapValues[y][x] = WALL
			case 'S':
				mapValues[y][x] = START
				startCoordinate = Coordinate{x, y}
			case 'E':
				endCoordinate = Coordinate{x, y}
				mapValues[y][x] = 1
			default:
				mapValues[y][x] = 1
			}
		}
	}

	// Build the linked list and calculate distances
	startElement := &RacetrackElement{coordinate: startCoordinate, next: nil}
	currentElement := startElement
	currentCoordinate := startCoordinate
	distance := 0
	visited := make(map[Coordinate]bool)
	visited[currentCoordinate] = true

	for currentCoordinate != endCoordinate {
		mapValues[currentCoordinate.Y][currentCoordinate.X] = distance

		// Find next coordinate
		directions := []Coordinate{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, dir := range directions {
			next := Coordinate{currentCoordinate.X + dir.X, currentCoordinate.Y + dir.Y}

			if next.Y >= 0 && next.Y < len(mapValues) &&
				next.X >= 0 && next.X < len(mapValues[next.Y]) &&
				mapValues[next.Y][next.X] != WALL &&
				!visited[next] {

				nextElement := &RacetrackElement{coordinate: next, next: nil}
				currentElement.next = nextElement
				currentElement = nextElement
				currentCoordinate = next
				visited[next] = true
				distance++
				break
			}
		}
	}

	// Set final distance for end coordinate
	mapValues[currentCoordinate.Y][currentCoordinate.X] = distance

	return RacetrackMap{
		values: mapValues,
		start:  *startElement,
		length: distance,
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
