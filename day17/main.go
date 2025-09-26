package day17

import (
	"fmt"
	"strings"
)

func ChronospatialComputerOutputFor(inputContent string) string {
	var computer = ParseChronospatialComputer(inputContent)
	computer.RunProgram()
	return joinSliceOfIntToString(computer.GetOutput())
}

func joinSliceOfIntToString(slice []int) string {
    return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice), " "), ","), "[]")
}
