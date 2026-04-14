package main

func intersect(nums1 []int, nums2 []int) []int {
	freqMap := map[int]int{}
	result := []int{}

	// Pass 1: count everything in nums1
	for _, v := range nums1 {
		freqMap[v]++
	}

	// Pass 2: walk nums2, claim tickets
	for _, v := range nums2 {
		if freqMap[v] > 0 {
			result = append(result, v)
			freqMap[v]--
		}
	}

	return result
}
