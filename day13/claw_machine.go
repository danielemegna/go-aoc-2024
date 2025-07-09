package day13

type ClawMachine struct {
	buttonA Button
	buttonB Button
	prizeCoordinate Coordinate
}

type Button = Coordinate

type Coordinate struct {
	X int
	Y int
}
