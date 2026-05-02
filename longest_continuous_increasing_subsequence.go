package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curLen++
		} else {
			curLen = 1
		}
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
