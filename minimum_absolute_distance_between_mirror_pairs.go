package main

import (
	"math"
)

func minMirrorPairDistance(nums []int) int {
	indexMap := make(map[int]int)
	smallest := math.MaxInt32

	for j, num := range nums {
		if i, exists := indexMap[num]; exists {
			diff := j - i
			if diff < smallest {
				smallest = diff
			}
		}
		reversed := reverse(num)
		indexMap[reversed] = j
	}

	if smallest == math.MaxInt32 {
		return -1
	}
	return smallest
}

func reverse(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}
