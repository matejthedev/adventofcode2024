package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type MaxHeap struct {
	slice    []int
	heapSize int
}

func (h *MaxHeap) heapify(length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && h.slice[left] > h.slice[largest] {
		largest = left
	}

	if right < length && h.slice[right] > h.slice[largest] {
		largest = right
	}

	if largest != i {
		h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
		h.heapify(length, largest)
	}
}

func (h *MaxHeap) sort() {
	length := len(h.slice)

	for i := length/2 - 1; i >= 0; i-- {
		h.heapify(length, i)
	}

	for i := length - 1; i >= 0; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapify(i, 0)
	}
}

//go:embed data.txt
var dataTxt string

func parseData() ([]int, []int, error) {
	var leftVals []int
	var rightVals []int

	lines := strings.Split(dataTxt, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %s", line)
		}

		leftNum, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number %s in line %s", fields[0], line)
		}
		rightNum, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number %s in line %s", fields[1], line)
		}

		leftVals = append(leftVals, leftNum)
		rightVals = append(rightVals, rightNum)
	}

	return leftVals, rightVals, nil
}
