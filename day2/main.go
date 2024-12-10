package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed data.txt
var dataTxt string

func parseData() [][]int {
	var reports [][]int

	lines := strings.Split(dataTxt, "\n")
	for _, line := range lines {
		numberStrings := strings.Fields(line)

		report := make([]int, len(numberStrings))
		for i, s := range numberStrings {
			report[i], _ = strconv.Atoi(s)
		}

		reports = append(reports, report)
	}

	return reports
}

func countSafeReports(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		if isReportSafe(report) {
			sum++
		}
	}
	return sum
}

func countSafeReportsWithTolerate(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		if isReportSafeWithTolerate(report) {
			sum++
		}
	}
	return sum
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	var prevDiff int
	directionEstablished := false
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		if !directionEstablished {
			prevDiff = diff
			directionEstablished = true
		} else {
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				return false
			}
			prevDiff = diff
		}
	}

	return true
}

func isReportSafeWithTolerate(report []int) bool {
	if isReportSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		newReport := append([]int(nil), report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}
