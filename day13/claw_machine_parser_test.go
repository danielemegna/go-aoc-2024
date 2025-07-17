package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSingleMachine(t *testing.T) {
	var machineStringRows = []string{
		"Button A: X+70, Y+10",
		"Button B: X+46, Y+19",
		"Prize: X=400, Y=500",
	}

	var machine = parseClawMachine(machineStringRows)

	var expected = ClawMachine{
		buttonA:         Coordinate{X: 70, Y: 10},
		buttonB:         Coordinate{X: 46, Y: 19},
		prizeCoordinate: Coordinate{X: 400, Y: 500},
	}
	assert.Equal(t, expected, machine)
}

func TestParseProvidedExampleMachines(t *testing.T) {
	var fileContent = SimulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)

	var machines = parseClawMachines(fileContent)

	assert.Len(t, machines, 4)
	assert.Equal(t,
		ClawMachine{
			buttonA:         Coordinate{X: 94, Y: 34},
			buttonB:         Coordinate{X: 22, Y: 67},
			prizeCoordinate: Coordinate{X: 8400, Y: 5400},
		},
		machines[0],
	)
	assert.Equal(t, 26, machines[1].buttonA.X)
	assert.Equal(t, 12176, machines[1].prizeCoordinate.Y)
	assert.Equal(t, 37, machines[2].buttonB.Y)
	assert.Equal(t,
		ClawMachine{
			buttonA:         Coordinate{X: 69, Y: 23},
			buttonB:         Coordinate{X: 27, Y: 71},
			prizeCoordinate: Coordinate{X: 18641, Y: 10279},
		},
		machines[3],
	)
}
