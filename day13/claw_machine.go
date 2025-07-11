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
	var buttonBPressCount = this.prizeCoordinate.X / this.buttonB.X

	for {
		var xPrizeCoordinateDiff = this.prizeCoordinate.X - (this.buttonB.X * buttonBPressCount)

		if xPrizeCoordinateDiff % this.buttonA.X == 0 {
			var buttonAPressCount = xPrizeCoordinateDiff / this.buttonA.X

			var y = (buttonAPressCount * this.buttonA.Y) + (buttonBPressCount * this.buttonB.Y)
			if y == this.prizeCoordinate.Y {
				return buttonAPressCount, buttonBPressCount
			}

			buttonBPressCount--
			if buttonBPressCount < 0 {
				break
			}

			continue
		}

		buttonBPressCount--
		if buttonBPressCount < 0 {
			break
		}
	}

	return -1, -1

}
