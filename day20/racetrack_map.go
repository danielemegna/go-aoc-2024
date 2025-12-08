package day20

import (
	"math"
	"strings"
)

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
const TRACK MapValue = math.MaxInt

func ParseRacetrack(rawMapString string) RacetrackMap {
	var mapValues, racetrackStartCoordinate = mapValuesParsing(rawMapString)
	var racetrackStart = RacetrackElement{Coordinate: racetrackStartCoordinate, Next: nil}

	var trackLength = 0
	var currentRacetrackElement = &racetrackStart
	for {
		var nextCoordinate *Coordinate
		for _, close := range currentRacetrackElement.Coordinate.CloseCoordinates() {
			if mapValues[close.Y][close.X] == TRACK {
				nextCoordinate = &close
				break
			}
		}

		if nextCoordinate == nil {
			break
		}

		mapValues[nextCoordinate.Y][nextCoordinate.X] = trackLength + 1
		var newRacetrackElement = RacetrackElement{Coordinate: *nextCoordinate}
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

func (this RacetrackMap) RacetrackStart() *RacetrackElement {
	return &this.start
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

func mapValuesParsing(rawMapString string) ([][]MapValue, Coordinate) {
	var inputLines = linesFrom(rawMapString)
	var rawMap = make([][]MapValue, len(inputLines))

	var racetrackStartCoordinate Coordinate
	for y := 0; y < len(inputLines); y++ {
		rawMap[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				rawMap[y][x] = WALL
			case 'S':
				rawMap[y][x] = START
				racetrackStartCoordinate = Coordinate{x, y}
			default:
				rawMap[y][x] = TRACK
			}
		}
	}
	return rawMap, racetrackStartCoordinate
}

func linesFrom(s string) []string {
	var trimmed = strings.TrimSpace(s)
	return strings.Split(trimmed, "\n")
}
