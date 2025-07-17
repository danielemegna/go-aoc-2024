package day13

import (
	"regexp"
	"strconv"
	"strings"
)

func parseClawMachines(fileContent string) []ClawMachine {
	var rows = rowsFrom(fileContent)
	var machines = []ClawMachine{}
	for i := 0; i < len(rows); i += 4 {
		machines = append(machines, parseClawMachine(rows[i:i+3]))
	}
	return machines
}

func parseClawMachine(rows []string) ClawMachine {
	var xButtonA, yButtonA = extractGroupsFrom(rows[0], `^Button A: X\+(\d+), Y\+(\d+)$`)
	var xButtonB, yButtonB = extractGroupsFrom(rows[1], `^Button B: X\+(\d+), Y\+(\d+)$`)
	var xPrize, yPrize = extractGroupsFrom(rows[2], `^Prize: X\=(\d+), Y\=(\d+)$`)

	return ClawMachine{
		buttonA:         Coordinate{X: xButtonA, Y: yButtonA},
		buttonB:         Coordinate{X: xButtonB, Y: yButtonB},
		prizeCoordinate: Coordinate{X: xPrize, Y: yPrize},
	}
}

func extractGroupsFrom(str string, regex string) (int, int) {
	var compiled, _ = regexp.Compile(regex)
	var matches = compiled.FindStringSubmatch(str)
	var firstValue, _ = strconv.Atoi(matches[1])
	var secondValue, _ = strconv.Atoi(matches[2])
	return firstValue, secondValue
}

func rowsFrom(s string) []string {
	var rows = strings.Split(s, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
