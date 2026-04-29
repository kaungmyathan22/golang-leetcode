package main

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}
	start := 0
	end := 0
	if len(nums) == 0 {
		return result
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i]+1 {
			end = i + 1
		} else {
			if start == end {
				result = append(result, fmt.Sprintf("%d", nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
			}
			start = i + 1
			end = i + 1
		}
	}

	// flush the last run  ← this was missing
	if start == end {
		result = append(result, fmt.Sprintf("%d", nums[start]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
	}

	return result
}
