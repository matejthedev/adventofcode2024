package main

import (
	_ "embed"
	"math"
)

func addUpDistances(s1, s2 []int) int {
	maxHeap1 := MaxHeap{s1, len(s1)}
	maxHeap2 := MaxHeap{s2, len(s2)}
	maxHeap1.sort()
	maxHeap2.sort()

	sum := 0
	for k, v := range s2 {
		sum = sum + int(math.Abs(float64(v)-float64(s1[k])))
	}
	return sum
}

func calculateSimilarityScore(s1, s2 []int) int {
	countInS2 := make(map[int]int)
	for _, num := range s2 {
		countInS2[num]++
	}

	score := 0
	for _, num := range s1 {
		count := countInS2[num]
		score += num * count
	}

	return score
}
