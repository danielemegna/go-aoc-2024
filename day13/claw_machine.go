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
	var buttonBNumerator = (this.prizeCoordinate.Y * this.buttonA.X) - (this.buttonA.Y * this.prizeCoordinate.X)
	var buttonBDenominator = (this.buttonB.Y * this.buttonA.X) - this.buttonB.X*this.buttonA.Y
	var buttonBPressCount = buttonBNumerator / buttonBDenominator

	if buttonBNumerator%buttonBDenominator != 0 {
		return -1, -1
	}

	var buttonAPressCount = (this.prizeCoordinate.X - (buttonBPressCount * this.buttonB.X)) / this.buttonA.X
	return buttonAPressCount, buttonBPressCount

}
