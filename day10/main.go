package main

import (
	"strings"
)

func getTrailHeadsTotalScore(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var explore func(r, c, h int, visited map[[2]int]bool)
	explore = func(r, c, h int, visited map[[2]int]bool) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return
		}
		if grid[r][c] != h {
			return
		}
		if h == 9 {
			visited[[2]int{r, c}] = true
			return
		}
		for _, d := range dirs {
			explore(r+d[0], c+d[1], h+1, visited)
		}
	}
	total := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				visited := make(map[[2]int]bool)
				explore(r, c, 0, visited)
				total += len(visited)
			}
		}
	}
	return total
}

func getTrailheadsRatingSum(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var explore func(r, c, h int) int
	explore = func(r, c, h int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return 0
		}
		if grid[r][c] != h {
			return 0
		}
		if h == 9 {
			return 1
		}
		sum := 0
		for _, d := range dirs {
			sum += explore(r+d[0], c+d[1], h+1)
		}
		return sum
	}
	ratingSum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				ratingSum += explore(r, c, 0)
			}
		}
	}
	return ratingSum
}
