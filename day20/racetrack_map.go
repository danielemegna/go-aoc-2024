package day20

import "strings"

type RacetrackMap = [][]MapValue
type MapValue = int

const (
	TRACK MapValue = iota
	WALL
)

type Racetrack struct {
	start RacetrackElement
}

type RacetrackElement struct {
	Coordinate        Coordinate
	DistanceFromStart int
	Next              *RacetrackElement
}

func ParseRacetrack(rawMapString string) Racetrack {
	var racetrackMap, racetrackStartCoordinate = mapParsing(rawMapString)
	var racetrackStartElement = RacetrackElement{
		Coordinate:        racetrackStartCoordinate,
		DistanceFromStart: 0,
		Next:              nil,
	}
	createRacetrackElementsChain(&racetrackStartElement, nil, racetrackMap)

	return Racetrack{
		start: racetrackStartElement,
	}
}

func (this Racetrack) RacetrackStart() *RacetrackElement {
	return &this.start
}

func mapParsing(rawMapString string) (RacetrackMap, Coordinate) {
	var inputLines = linesFrom(rawMapString)
	var racetrackMap = make(RacetrackMap, len(inputLines))

	var racetrackStartCoordinate Coordinate
	for y := 0; y < len(inputLines); y++ {
		racetrackMap[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				racetrackMap[y][x] = WALL
			case 'S':
				racetrackStartCoordinate = Coordinate{x, y}
			default:
				racetrackMap[y][x] = TRACK
			}
		}
	}
	return racetrackMap, racetrackStartCoordinate
}

func createRacetrackElementsChain(current *RacetrackElement, previous *RacetrackElement, mapValues RacetrackMap) {
	var nextCoordinate = findNextRacetrackCoordinate(current.Coordinate, previous, mapValues)
	if nextCoordinate == nil {
		return
	}

	var newRacetrackElement = RacetrackElement{
		Coordinate:        *nextCoordinate,
		DistanceFromStart: current.DistanceFromStart + 1,
		Next:              nil,
	}
	current.Next = &newRacetrackElement
	createRacetrackElementsChain(&newRacetrackElement, current, mapValues)
}

func findNextRacetrackCoordinate(current Coordinate, previous *RacetrackElement, mapValues RacetrackMap) *Coordinate {
	for _, close := range current.CloseCoordinates() {
		if mapValues[close.Y][close.X] == WALL {
			continue
		}
		if previous != nil && close == previous.Coordinate {
			continue
		}

		return &close
	}

	return nil
}

func linesFrom(s string) []string {
	var trimmed = strings.TrimSpace(s)
	return strings.Split(trimmed, "\n")
}
