package day15

import (
	"strings"
)

// WarehouseParser parses the input file to extract the warehouse map and the robot's moves.
type WarehouseParser struct{}

// NewWarehouseParser creates a new warehouse parser.
func NewWarehouseParser() *WarehouseParser {
	return &WarehouseParser{}
}

// Parse parses the input file content and returns the warehouse map and the robot's moves.
func (p *WarehouseParser) Parse(content string) (*WarehouseMap, []rune) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	
	// Find the empty line that separates the map from the moves
	emptyLineIndex := -1
	for i, line := range lines {
		if line == "" {
			emptyLineIndex = i
			break
		}
	}
	
	if emptyLineIndex == -1 {
		panic("Invalid input format: no empty line found to separate map from moves")
	}
	
	// Parse the map
	mapLines := lines[:emptyLineIndex]
	warehouseMap := p.parseMap(mapLines)
	
	// Parse the moves
	movesLines := lines[emptyLineIndex+1:]
	moves := p.parseMoves(movesLines)
	
	return warehouseMap, moves
}

// parseMap parses the map part of the input and creates a WarehouseMap.
func (p *WarehouseParser) parseMap(mapLines []string) *WarehouseMap {
	warehouseMap := NewWarehouseMap()
	
	for _, line := range mapLines {
		warehouseMap.AddRow(line)
	}
	
	return warehouseMap
}

// parseMoves parses the moves part of the input and returns a slice of runes.
func (p *WarehouseParser) parseMoves(movesLines []string) []rune {
	// Join all moves lines into a single string and convert to runes
	movesStr := strings.Join(movesLines, "")
	return []rune(movesStr)
}