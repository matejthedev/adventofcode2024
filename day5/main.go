package main

import (
	"strconv"
	"strings"
)

func getMiddlePageNumSum(input string) int {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	updates, rules := getUpdatesAndRules(rows)

	for _, currentUpdate := range updates {
		isValid := true
		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}
		}

		if isValid {
			middle := currentUpdate[len(currentUpdate)/2]
			output += middle
		}
	}

	return output
}

func getOrderedMiddlePageNumSum(input string) int {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	updates, rules := getUpdatesAndRules(rows)

	for _, currentUpdate := range updates {
		isValid := true
		for i := len(currentUpdate) - 1; i > 0; i-- {
			currentNum := currentUpdate[i]
			for _, num := range currentUpdate[:i] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}
		}

		if !isValid {
			var ordered []int
			remaining := map[int]bool{}
			dependencies := make(map[int][]int)

			for _, num := range currentUpdate {
				remaining[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(remaining) > 0 {
				for num := range remaining {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(remaining, num)

						for key, val := range dependencies {
							var newList []int

							for _, n := range val {
								if n != num {
									newList = append(newList, n)
								}
							}

							dependencies[key] = newList
						}
					}
				}
			}

			middle := ordered[len(ordered)/2]
			output += middle
		}
	}

	return output
}

func getUpdatesAndRules(rows []string) ([][]int, map[int][]int) {
	rules := map[int][]int{}
	isAddRule := true
	var updates [][]int

	for _, line := range rows {
		if len(line) == 0 {
			isAddRule = false
			continue
		}

		if isAddRule {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			rules[num1] = append(rules[num1], num2)
		} else {
			var currentUpdate []int
			parts := strings.Split(line, ",")

			for _, part := range parts {
				num1, _ := strconv.Atoi(part)
				currentUpdate = append(currentUpdate, num1)
			}

			updates = append(updates, currentUpdate)
		}
	}
	return updates, rules
}
