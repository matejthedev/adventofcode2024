package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getTotalCalibrationResult(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))

		for i, numString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numString)
		}

		if evaluateCalc(target, numbers, 0, 0) {
			result += target
		}
	}

	return result
}

func getTotalResultWithInsertingOperators(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))

		for i, numString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numString)
		}

		if evaluateWithOperators(target, numbers, 0, 0) {
			result += target
		}
	}

	return result
}

func evaluateCalc(target int, numbers []int, index int, currentResult int) bool {
	if index >= len(numbers) {
		return currentResult == target
	}

	currentNumber := numbers[index]

	newResult := currentResult + currentNumber
	if newResult <= target {
		if evaluateCalc(target, numbers, index+1, newResult) {
			return true
		}
	}

	newResult = currentResult * currentNumber
	if newResult <= target {
		return evaluateCalc(target, numbers, index+1, newResult)
	}

	return false
}

func evaluateWithOperators(target int, numbers []int, index int, currentResult int) bool {
	if index >= len(numbers) {
		return currentResult == target
	}

	currentNumber := numbers[index]

	newResult := currentResult + currentNumber
	if newResult <= target {
		if evaluateWithOperators(target, numbers, index+1, newResult) {
			return true
		}
	}

	newResult = currentResult * currentNumber
	if newResult <= target {
		if evaluateWithOperators(target, numbers, index+1, newResult) {
			return true
		}
	}

	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, currentNumber))
	if concatenated <= target {
		correct := evaluateWithOperators(target, numbers, index+1, concatenated)
		return correct
	}

	return false
}
