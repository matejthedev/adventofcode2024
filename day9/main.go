package main

import (
	"strconv"
)

func parseBlocks(input string) ([]rune, int) {
	res := []rune{}
	isFile := true
	fileNum := 0
	for _, ch := range input {
		n, _ := strconv.Atoi(string(ch))
		if isFile {
			for i := 0; i < n; i++ {
				res = append(res, rune('0'+fileNum))
			}
			fileNum++
		} else {
			for i := 0; i < n; i++ {
				res = append(res, '.')
			}
		}
		isFile = !isFile
	}
	return res, fileNum
}

func getFsChecksum(input string) int {
	blocks, _ := parseBlocks(input)
	emptyIdx := 0
	lastIdx := len(blocks) - 1
	for blocks[emptyIdx] != '.' {
		emptyIdx++
	}
	for blocks[lastIdx] == '.' {
		lastIdx--
	}
	for emptyIdx <= lastIdx {
		blocks[emptyIdx], blocks[lastIdx] = blocks[lastIdx], '.'
		for blocks[emptyIdx] != '.' {
			emptyIdx++
		}
		for blocks[lastIdx] == '.' {
			lastIdx--
		}
	}
	blockIdx := 0
	total := 0
	for blocks[blockIdx] != '.' {
		total += blockIdx * int(blocks[blockIdx]-'0')
		blockIdx++
	}
	return total
}

func getCompactFsChecksum(input string) int {
	blocks, fileCount := parseBlocks(input)
	for currentFile := fileCount - 1; currentFile >= 0; currentFile-- {
		filePositions := []int{}
		for i, block := range blocks {
			if block == rune('0'+currentFile) {
				filePositions = append(filePositions, i)
			}
		}
		startIdx := -1
		length := 0
		for i := 0; i < filePositions[0]; i++ {
			if blocks[i] != '.' {
				startIdx = -1
				length = 0
				continue
			}
			if startIdx == -1 {
				startIdx = i
			}
			length++
			if length == len(filePositions) {
				break
			}
		}
		if length == len(filePositions) {
			for i := 0; i < length; i++ {
				blocks[filePositions[i]], blocks[startIdx+i] = '.', rune('0'+currentFile)
			}
		}
	}
	total := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != '.' {
			total += i * int(blocks[i]-'0')
		}
	}
	return total
}
