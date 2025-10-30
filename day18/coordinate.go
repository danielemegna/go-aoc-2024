package day18

import (
	"regexp"
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

func ParseCoordinateFrom(raw string) Coordinate {
	var r, _ = regexp.Compile(`(\d+),(\d+)`)
	var matches = r.FindStringSubmatch(raw)
	var X, _ = strconv.Atoi(matches[1])
	var Y, _ = strconv.Atoi(matches[2])
	return Coordinate{X, Y}
}

func (c Coordinate) CloseCoordinates() []Coordinate {
	return []Coordinate{
		{c.X - 1, c.Y},
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
	}
}

type CoordinateInPath struct {
	coordinate Coordinate
	parent     *CoordinateInPath
	pathLength int
}

func LinkCoordinateInPath(coordinate Coordinate, parent *CoordinateInPath) CoordinateInPath {
	return CoordinateInPath{
		coordinate: coordinate,
		parent:     parent,
		pathLength: parent.pathLength + 1,
	}
}

func (this CoordinateInPath) PathCoordinates() []Coordinate {
	var currentNode *CoordinateInPath = &this
	var result = []Coordinate{}
	for currentNode != nil {
		result = append(result, currentNode.coordinate)
		currentNode = currentNode.parent
	}
	return result
}
