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

	var guessBPressCount = this.prizeCoordinate.X / this.buttonB.X

	for {
		var differenza = this.prizeCoordinate.X - (this.buttonB.X * guessBPressCount)

		var resto = differenza % this.buttonA.X
		if resto == 0 {
			var buttonAPressCount = differenza / this.buttonA.X
			var buttonBPressCount = guessBPressCount

			var y = (buttonAPressCount * this.buttonA.Y) + (buttonBPressCount * this.buttonB.Y)
			if y == this.prizeCoordinate.Y {
				return buttonAPressCount, buttonBPressCount
			}

			guessBPressCount--
			if guessBPressCount < 0 {
				break
			}

			continue
		}

		guessBPressCount--
		if guessBPressCount < 0 {
			break
		}
	}

	return -1, -1

}
