package main

import (
	"strings"
)

const (
	UP    = '^'
	RIGHT = '>'
	DOWN  = 'v'
	LEFT  = '<'
)

func getVisitedPositions(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	visited := 1

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	row, column := mapPosition(grid)

	for 0 < column && column < (len(lines[0])-1) && 0 < row && row < (len(lines)-1) {

		if grid[row][column] == UP {
			if grid[row-1][column] == '#' {
				grid[row][column] = RIGHT
				continue
			} else {
				grid[row][column] = 'x'
				if grid[row-1][column] != 'x' {
					visited += 1
				}
				grid[row-1][column] = UP
				row -= 1
				continue
			}
		}

		if grid[row][column] == RIGHT {
			if grid[row][column+1] == '#' {
				grid[row][column] = DOWN
				continue
			} else {
				grid[row][column] = 'x'
				if grid[row][column+1] != 'x' {
					visited += 1
				}
				grid[row][column+1] = RIGHT
				column += 1
				continue
			}
		}

		if grid[row][column] == DOWN {
			if grid[row+1][column] == '#' {
				grid[row][column] = LEFT
				continue
			} else {
				grid[row][column] = 'x'
				if grid[row+1][column] != 'x' {
					visited += 1
				}
				grid[row+1][column] = DOWN
				row += 1
				continue
			}
		}

		if grid[row][column] == LEFT {
			if grid[row][column-1] == '#' {
				grid[row][column] = UP
				continue
			} else {
				grid[row][column] = 'x'
				if grid[row][column-1] != 'x' {
					visited += 1
				}
				grid[row][column-1] = LEFT
				column -= 1
				continue
			}
		}
	}

	return visited
}

func mapPosition(input [][]rune) (int, int) {
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			if input[row][column] == '^' {
				return row, column
			}
			if input[row][column] == '>' {
				return row, column
			}
			if input[row][column] == 'v' {
				return row, column
			}
			if input[row][column] == '<' {
				return row, column
			}
		}
	}
	return 0, 0
}

func getNumberOfObstructions(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}
	obstructions := 0

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for row, _ := range grid {
		if guardRow >= 0 {
			break
		}
		for col, _ := range grid[row] {
			if grid[row][col] == '^' {
				guardRow = row
				guardCol = col
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				continue
			}

			grid[row][col] = '#'

			visitedLocations := make(map[[3]int]bool)
			currentRow := guardRow
			currentCol := guardCol
			currentDirection := guardDirection

			stuckInALoop := false

			for {
				guardState := [3]int{currentRow, currentCol, currentDirection}
				if visitedLocations[guardState] {
					stuckInALoop = true
					break
				}

				visitedLocations[guardState] = true

				nextGuardRow := currentRow + directions[currentDirection][0]
				nextGuardCol := currentCol + directions[currentDirection][1]

				if nextGuardRow < 0 || nextGuardRow >= len(grid) || nextGuardCol < 0 || nextGuardCol >= len(grid[0]) {
					break
				}

				if grid[nextGuardRow][nextGuardCol] == '#' {
					currentDirection = (currentDirection + 1) % 4
				} else {
					currentRow = nextGuardRow
					currentCol = nextGuardCol
				}
			}

			if stuckInALoop {
				obstructions++
			}

			grid[row][col] = '.'
		}
	}

	return obstructions
}
