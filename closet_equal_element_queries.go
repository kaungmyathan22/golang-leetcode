package main

import "sort"

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	// Step 1: group indices
	posMap := make(map[int][]int)
	for i, v := range nums {
		posMap[v] = append(posMap[v], i)
	}

	res := make([]int, len(queries))

	for qi, q := range queries {
		positions := posMap[nums[q]]

		if len(positions) == 1 {
			res[qi] = -1
			continue
		}

		// find index in sorted list
		k := sort.SearchInts(positions, q)

		// circular neighbors
		left := positions[(k-1+len(positions))%len(positions)]
		right := positions[(k+1)%len(positions)]

		// distances
		d1 := abs(q - left)
		d2 := abs(q - right)

		d1 = min(d1, n-d1)
		d2 = min(d2, n-d2)

		res[qi] = min(d1, d2)
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
