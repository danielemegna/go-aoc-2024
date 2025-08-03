package day14

import (
	"strings"
	"time"
)

func SafetyFactorAfter100Seconds(fileContent string, spaceSize SpaceSize) int {
	var robots = ParseRobotLines(fileContent)
	var space = Space{size: spaceSize, guards: robots}

	space.AfterSeconds(100)

	return space.GetSafetyFactor()
}

// run this with `go run day14/run/run.go`
func FindTree() {
	var space = buildSpace()

	for {
		space.Print()
		space.AfterSeconds(1)
		time.Sleep(500 * time.Millisecond)
	}

}

func buildSpace() Space {
	var providedExampleInputLines = []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	var fileContent = strings.Join(providedExampleInputLines, "\n") + "\n"
	var spaceSize = SpaceSize{width: 11, height: 7}
	var robots = ParseRobotLines(fileContent)
	return Space{size: spaceSize, guards: robots}
}
