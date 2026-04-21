package main

func MinimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)

	// DSU / Union-Find
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // path compression
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa := find(a)
		pb := find(b)
		if pa != pb {
			parent[pb] = pa
		}
	}

	// Step 1: union all swaps
	for _, pair := range allowedSwaps {
		union(pair[0], pair[1])
	}

	// Step 2: group indices by root
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := find(i)
		groups[root] = append(groups[root], i)
	}

	distance := 0

	// Step 3: process each group
	for _, indices := range groups {
		count := make(map[int]int)

		// count source values
		for _, i := range indices {
			count[source[i]]++
		}

		// match with target
		for _, i := range indices {
			if count[target[i]] > 0 {
				count[target[i]]--
			} else {
				distance++
			}
		}
	}

	return distance
}
