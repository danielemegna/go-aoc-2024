package day20

import "strings"

func CountCheatsSavingAtLeast(picoseconds int, fileContent string) int {
	lines := strings.Split(strings.TrimSpace(fileContent), "\n")
	racetrackMap := ParseRacetrack(lines)
	cheats := PossibleCheatsIn(racetrackMap)

	count := 0
	for _, cheat := range cheats {
		if cheat.savingInPicoseconds >= picoseconds {
			count++
		}
	}

	return count
}
