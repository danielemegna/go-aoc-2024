package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClawMachineWithoutPrizeSolution(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 26, Y: 66},
		buttonB:         Coordinate{X: 67, Y: 21},
		prizeCoordinate: Coordinate{X: 12748, Y: 12176},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, -1, buttonAPressCount)
	assert.Equal(t, -1, buttonBPressCount)
}

func TestClawMachineWithPrizeSolution(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 94, Y: 34},
		buttonB:         Coordinate{X: 22, Y: 67},
		prizeCoordinate: Coordinate{X: 8400, Y: 5400},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, 80, buttonAPressCount)
	assert.Equal(t, 40, buttonBPressCount)
}

func TestAnotherClawMachineWithPrizeSolution(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 17, Y: 86},
		buttonB:         Coordinate{X: 84, Y: 37},
		prizeCoordinate: Coordinate{X: 7870, Y: 6450},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, 38, buttonAPressCount)
	assert.Equal(t, 86, buttonBPressCount)
}

func TestAnotherClawMachineWithoutPrizeSolution(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 69, Y: 23},
		buttonB:         Coordinate{X: 27, Y: 71},
		prizeCoordinate: Coordinate{X: 18641, Y: 10279},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, -1, buttonAPressCount)
	assert.Equal(t, -1, buttonBPressCount)
}

func TestClawMachineWithPrizeVeryFar(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 26, Y: 66},
		buttonB:         Coordinate{X: 67, Y: 21},
		prizeCoordinate: Coordinate{X: 10000000012748, Y: 10000000012176},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, 118679050709, buttonAPressCount)
	assert.Equal(t, 103199174542, buttonBPressCount)
}

func TestClawMachineWithoutSolutionWithPrizeVeryFar(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 94, Y: 34},
		buttonB:         Coordinate{X: 22, Y: 67},
		prizeCoordinate: Coordinate{X: 10000000008400, Y: 10000000005400},
	}

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, -1, buttonAPressCount)
	assert.Equal(t, -1, buttonBPressCount)
}

func TestProvidedExampleClawMachineWithFixedPrizeCoordinates(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 26, Y: 66},
		buttonB:         Coordinate{X: 67, Y: 21},
		prizeCoordinate: Coordinate{X: 12748, Y: 12176},
	}

	machine.FixPrizeCoordinates()

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, 118679050709, buttonAPressCount)
	assert.Equal(t, 103199174542, buttonBPressCount)
}

func TestAnotherProvidedExampleClawMachineWithFixedPrizeCoordinates(t *testing.T) {
	var machine = ClawMachine{
		buttonA:         Coordinate{X: 69, Y: 23},
		buttonB:         Coordinate{X: 27, Y: 71},
		prizeCoordinate: Coordinate{X: 18641, Y: 10279},
	}

	machine.FixPrizeCoordinates()

	var buttonAPressCount, buttonBPressCount = machine.HowToWinThePrize()

	assert.Equal(t, 102851800151, buttonAPressCount)
	assert.Equal(t, 107526881786, buttonBPressCount)
}
