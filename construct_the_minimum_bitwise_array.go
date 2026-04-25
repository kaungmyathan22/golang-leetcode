package main

func minBitwiseArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		found := false
		// Try all values from 0 up to nums[i]-1
		for x := 0; x < nums[i]; x++ {
			if (x | (x + 1)) == nums[i] {
				res[i] = x
				found = true
				break
			}
		}
		if !found {
			res[i] = -1
		}
	}
	return res
}
