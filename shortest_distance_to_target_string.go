package main

func ClosestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	minDist := -1

	for i := 0; i < n; i++ {
		if words[i] == target {
			dist := Abs(i - startIndex)
			// For circular array, take the minimum distance going forward and backward
			circDist := min(dist, n-dist)
			if minDist == -1 || circDist < minDist {
				minDist = circDist
			}
		}
	}

	return minDist
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
