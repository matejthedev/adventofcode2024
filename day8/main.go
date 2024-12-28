package main

import (
	"strings"
)

func countAntiNodes(input string) int {
	grid := parseGrid(input)
	antennas := collectAntennaPositions(grid)

	antiNodes := make(map[[2]int]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				deltaRow := pos2[0] - pos1[0]
				deltaCol := pos2[1] - pos1[1]

				antiNode1 := [2]int{pos1[0] - deltaRow, pos1[1] - deltaCol}
				antiNode2 := [2]int{pos2[0] + deltaRow, pos2[1] + deltaCol}

				if inBounds(antiNode1[0], antiNode1[1], len(grid), len(grid[0])) {
					antiNodes[antiNode1] = true
				}
				if inBounds(antiNode2[0], antiNode2[1], len(grid), len(grid[0])) {
					antiNodes[antiNode2] = true
				}
			}
		}
	}

	return len(antiNodes)
}

func countUniqueLocations(input string) int {
	grid := parseGrid(input)
	antennas := collectAntennaPositions(grid)

	antiNodes := make(map[[2]int]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				deltaRow := pos2[0] - pos1[0]
				deltaCol := pos2[1] - pos1[1]

				r1, c1 := pos1[0], pos1[1]
				for inBounds(r1, c1, len(grid), len(grid[0])) {
					antiNodes[[2]int{r1, c1}] = true
					r1 -= deltaRow
					c1 -= deltaCol
				}

				r2, c2 := pos2[0], pos2[1]
				for inBounds(r2, c2, len(grid), len(grid[0])) {
					antiNodes[[2]int{r2, c2}] = true
					r2 += deltaRow
					c2 += deltaCol
				}
			}
		}
	}

	return len(antiNodes)
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func collectAntennaPositions(grid [][]rune) map[rune][][2]int {
	antennas := make(map[rune][][2]int)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if ch := grid[row][col]; ch != '.' {
				antennas[ch] = append(antennas[ch], [2]int{row, col})
			}
		}
	}
	return antennas
}

func inBounds(r, c, maxRow, maxCol int) bool {
	return r >= 0 && r < maxRow && c >= 0 && c < maxCol
}
