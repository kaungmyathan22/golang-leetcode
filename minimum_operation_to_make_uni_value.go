package main

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])
	nums := make([]int, 0, m*n)
	base := grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (grid[i][j]-base)%x != 0 {
				return -1
			}
			nums = append(nums, grid[i][j])
		}
	}

	sort.Ints(nums)
	median := nums[len(nums)/2]
	fmt.Println(nums)
	fmt.Println(median)
	ops := 0
	for _, num := range nums {
		ops += abs(num-median) / x
	}
	return ops
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
