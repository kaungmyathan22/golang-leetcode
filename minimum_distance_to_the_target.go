package main

func getMinDistance(nums []int, target int, start int) int {
	minDist := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			dist := abs(i - start)
			if dist < minDist {
				minDist = dist
			}
		}
	}
	if minDist == len(nums) {
		return -1
	}
	return minDist
}
