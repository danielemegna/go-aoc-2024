package day17

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseChronospatialComputer(fileContent string) ChronospatialComputer {
	var lines = linesFrom(fileContent)
	var registerA = parseRegister(lines[0], 'A')
	var registerB = parseRegister(lines[1], 'B')
	var registerC = parseRegister(lines[2], 'C')
	var instructions = parseProgram(lines[4])
	
	return NewChronospatialComputer(registerA, registerB, registerC, instructions)
}

func parseRegister(line string, register rune) int {
	var compiled, _ = regexp.Compile(`^Register ` + string(register) + `: (\d+)$`)
	var matches = compiled.FindStringSubmatch(line)
	var value, _ = strconv.Atoi(matches[1])
	return value
}

func parseProgram(line string) []int {
	var compiled, _ = regexp.Compile(`^Program: (.*)$`)
	var matches = compiled.FindStringSubmatch(line)
	var programString = matches[1]
	var instructionStrings = strings.Split(programString, ",")
	
	var instructions = []int{}
	for _, instructionString := range instructionStrings {
		var instruction, _ = strconv.Atoi(instructionString)
		instructions = append(instructions, instruction)
	}
	
	return instructions
}

func linesFrom(s string) []string {
	var rows = strings.Split(s, "\n")
	rows = rows[:len(rows)-1]
	return rows
}