package main

func canBeEqual(s1 string, s2 string) bool {
	if len(s1) != 4 || len(s2) != 4 {
		return false
	}

	group1s1 := []byte{s1[0], s1[2]}
	group1s2 := []byte{s2[0], s2[2]}
	group2s1 := []byte{s1[1], s1[3]}
	group2s2 := []byte{s2[1], s2[3]}

	// Sort the groups for comparison
	if group1s1[0] > group1s1[1] {
		group1s1[0], group1s1[1] = group1s1[1], group1s1[0]
	}
	if group1s2[0] > group1s2[1] {
		group1s2[0], group1s2[1] = group1s2[1], group1s2[0]
	}
	if group2s1[0] > group2s1[1] {
		group2s1[0], group2s1[1] = group2s1[1], group2s1[0]
	}
	if group2s2[0] > group2s2[1] {
		group2s2[0], group2s2[1] = group2s2[1], group2s2[0]
	}

	return group1s1[0] == group1s2[0] && group1s1[1] == group1s2[1] &&
		group2s1[0] == group2s2[0] && group2s1[1] == group2s2[1]
}
