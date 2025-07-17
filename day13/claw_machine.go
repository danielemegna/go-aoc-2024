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
	if buttonBNumerator%buttonBDenominator != 0 {
		return -1, -1
	}

	var buttonBPressCount = buttonBNumerator / buttonBDenominator

	var buttonANumerator = this.prizeCoordinate.X - (buttonBPressCount * this.buttonB.X)
	var buttonADenominator = this.buttonA.X
	if buttonANumerator % buttonADenominator != 0 {
		return -1, -1
	}

	var buttonAPressCount = buttonANumerator / buttonADenominator

	return buttonAPressCount, buttonBPressCount
}

func (this *ClawMachine) FixPrizeCoordinates() {
	this.prizeCoordinate.X += 10000000000000
	this.prizeCoordinate.Y += 10000000000000
}
