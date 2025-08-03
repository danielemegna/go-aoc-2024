package day14

import (
	"fmt"
	"os"
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
	var elapsedSeconds = 0
	var patternRepeatLimit = 10403

	for {
		space.AfterSeconds(1)
		elapsedSeconds++
		var shouldPrint = space.GetSafetyFactorWith(40) == 0

		if shouldPrint {
			space.Print()
			fmt.Printf("======== Elapsed seconds: %d\n", elapsedSeconds)
			time.Sleep(200 * time.Millisecond)
		}

		if elapsedSeconds > patternRepeatLimit {
			break
		}
	}

}

func buildSpace() Space {
	var fileBytes, _ = os.ReadFile("day14/input.txt")
	var fileContent = string(fileBytes)
	var spaceSize = SpaceSize{width: 101, height: 103}
	var robots = ParseRobotLines(fileContent)
	return Space{size: spaceSize, guards: robots}
}
