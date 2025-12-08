package day20

import "strings"

type RacetrackMap struct {
	values [][]MapValue
	start  RacetrackElement
	length int
}

type RacetrackElement struct {
	Coordinate Coordinate
	Next       *RacetrackElement
}

type MapValue = int

const START MapValue = 0
const WALL MapValue = -1

func ParseRacetrack(rawMap string) RacetrackMap {
	var inputLines = linesFrom(rawMap)
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
				racetrackStart.Coordinate = Coordinate{x, y}
			default:
				mapValues[y][x] = 1
			}
		}
	}

	var trackLength = 0
	var currentRacetrackElement = &racetrackStart
	for {
		var c = currentRacetrackElement.Coordinate
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

		mapValues[nextCoordinate.Y][nextCoordinate.X] = trackLength + 1
		var newRacetrackElement = RacetrackElement{Coordinate: nextCoordinate}
		currentRacetrackElement.Next = &newRacetrackElement
		currentRacetrackElement = currentRacetrackElement.Next
		trackLength++
	}

	return RacetrackMap{
		values: mapValues,
		start:  racetrackStart,
		length: trackLength,
	}
}

func (this RacetrackMap) RacetrackStart() RacetrackElement {
	return this.start
}

func (this RacetrackMap) ValueAt(c Coordinate) MapValue {
	return this.values[c.Y][c.X]
}

func (this RacetrackMap) RacetrackLength() int {
	return this.length
}

func (this RacetrackMap) MapSize() int {
	return len(this.values) // assuming always square maps
}

func linesFrom(s string) []string {
	var trimmed = strings.TrimSpace(s)
	return strings.Split(trimmed, "\n")
}
