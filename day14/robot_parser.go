package day14

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

func ParseRobotLines(robotLinesAsString string) []RobotGuard {
	var lines = linesFrom(robotLinesAsString)
	return lo.Map(lines, func(line string, _ int) RobotGuard {
		return ParseRobotGuard(line)
	})
}

func ParseRobotGuard(robotAsString string) RobotGuard {
	var positionX, positionY, horizontalVelocity, verticalVelocity = extractRobotDataFrom(robotAsString)
	return RobotGuard{
		position: Position{positionX, positionY},
		velocity: Velocity{horizontalVelocity, verticalVelocity},
	}
}

func extractRobotDataFrom(str string) (int, int, int, int) {
	var compiled, _ = regexp.Compile(`^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`)
	var matches = compiled.FindStringSubmatch(str)
	var positionX, _ = strconv.Atoi(matches[1])
	var positionY, _ = strconv.Atoi(matches[2])
	var horizontalVelocity, _ = strconv.Atoi(matches[3])
	var verticalVelocity, _ = strconv.Atoi(matches[4])
	return positionX, positionY, horizontalVelocity, verticalVelocity
}

func linesFrom(s string) []string {
	var rows = strings.Split(s, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
