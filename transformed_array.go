package main

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			res[i] = nums[i]
			continue
		}
		idx := ((i+nums[i])%n + n) % n
		res[i] = nums[idx]
	}
	return res
}
