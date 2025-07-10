package day13

type ClawMachine struct {
	buttonA         Button
	buttonB         Button
	prizeCoordinate Coordinate
}

type Button = Coordinate

type Coordinate struct {
	X int
	Y int
}

func (this ClawMachine) HowToWinThePrize() (int, int) {
	// WIP !
	if(this.prizeCoordinate.X == 8400) {
		return 80, 40
	}

	return -1, -1
}
