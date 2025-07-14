package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSingleMachine(t *testing.T) {
	var machineStringContent = SimulateFileContent([]string{
		"Button A: X+70, Y+10",
		"Button B: X+46, Y+19",
		"Prize: X=400, Y=500",
	})
	var machine = parseClawMachine(machineStringContent)
	var expected = ClawMachine{
		buttonA:         Coordinate{X: 70, Y: 10},
		buttonB:         Coordinate{X: 46, Y: 19},
		prizeCoordinate: Coordinate{X: 400, Y: 500},
	}
	assert.Equal(t, expected, machine)
}

func TestParseProvidedExampleMachines(t *testing.T) {
	t.Skip("WIP")
	var fileContent = SimulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	var machines = parseClawMachines(fileContent)
	assert.Len(t, machines, 4)
}
