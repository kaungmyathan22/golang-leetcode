package main

func distance(nums []int) []int64 {
	n := len(nums)
	res := make([]int64, n)

	// Map from value to indices [1, 3, 1, 1, 2] -> {1: [0, 2, 3], 3: [1], 2: [4]}
	indicesMap := make(map[int][]int)
	for i, v := range nums {
		indicesMap[v] = append(indicesMap[v], i)
	}
	for _, indices := range indicesMap {
		m := len(indices)
		if m == 1 {
			res[indices[0]] = 0
			continue
		}

		// Prefix sums of indices
		prefix := make([]int64, m+1)
		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(indices[i])
		}

		for i := 0; i < m; i++ {
			leftCnt := int64(i)
			leftSum := int64(indices[i])*leftCnt - prefix[i]
			rightCnt := int64(m - i - 1)
			rightSum := (prefix[m] - prefix[i+1]) - int64(indices[i])*rightCnt
			res[indices[i]] = leftSum + rightSum
		}
	}

	return res
}
