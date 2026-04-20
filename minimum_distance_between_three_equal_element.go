package main

import "fmt"

func MinimumDistanceBetweenThreeEqualElements(nums []int) int {
	positions := make(map[int][]int)
	for i, v := range nums {
		positions[v] = append(positions[v], i)
	}
	fmt.Println(positions)
	minDist := -1
	for _, indices := range positions {
		if len(indices) >= 3 {
			// Slide window of size 3
			for i := 0; i <= len(indices)-3; i++ {
				span := indices[i+2] - indices[i]
				dist := 2 * span
				if minDist == -1 || dist < minDist {
					minDist = dist
				}
			}
		}
	}
	return minDist
}
