package main

func minimumPairRemoval(nums []int) int {
	ops := 0

	isSorted := func() bool {
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				return false
			}
		}
		return true
	}

	for len(nums) > 1 && !isSorted() {
		minSum := nums[0] + nums[1]
		idx := 0
		for i := 1; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if sum < minSum {
				minSum = sum
				idx = i
			}
		}
		// Replace leftmost such minimal pair
		newNums := append([]int{}, nums[:idx]...)
		newNums = append(newNums, minSum)
		newNums = append(newNums, nums[idx+2:]...)
		nums = newNums
		ops++
	}
	return ops
}
