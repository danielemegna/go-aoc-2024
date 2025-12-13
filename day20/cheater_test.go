package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPossibileCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	var expected = []Cheat{
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 8 - 2 - 2}, // 8 start, 2 end, 2 cheat duration
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 7 - 3 - 2}, // 7 start, 3 end, 2 cheat duration
	}
	assert.ElementsMatch(t, expected, PossibleCheatsIn(racetrackMap, 2))
}

func TestFindPossibileCheatsInAnotherTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(ANOTHER_RACETRACK_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 2), 6)
}

func TestFindPossibileCheatsInProvidedExampleTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 2), 14 + 14 + 2 + 4 + 2 + 3 + 1 + 1 + 1 + 1 + 1)
}

func TestFindPossibileThreeStepCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	var expected = []Cheat{
		{startPosition: Coordinate{1, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 8 - 1 - 3}, // 8 start, 1 end, 3 cheat duration
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 8 - 2 - 2}, // 8 start, 2 end, 2 cheat duration
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 4}, savingInPicoseconds: 9 - 2 - 3}, // 9 start, 2 end, 3 cheat duration
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 7 - 2 - 3}, // 7 start, 2 end, 3 cheat duration
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 7 - 3 - 2}, 
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 8 - 3 - 3}, 
	}
	assert.ElementsMatch(t, expected, PossibleCheatsIn(racetrackMap, 3))
}

func TestFindPossibileThreeStepCheatsInAnotherTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(ANOTHER_RACETRACK_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 3), 16)
}

func TestFindPossibileLongCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 20), 9)
}

func TestFindPossibileLongCheatsInAnotherTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(ANOTHER_RACETRACK_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 20), 44)
}

func TestFindPossibileLongCheatsInProvidedExampleTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))
	assert.Len(t, PossibleCheatsIn(racetrackMap, 20), 3081)
}
