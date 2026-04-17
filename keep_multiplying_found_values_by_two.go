package main

func findFinalValue(nums []int, original int) int {
	for {
		found := false
		for _, num := range nums {
			if num == original {
				original *= 2
				found = true
			}
		}
		if !found {
			break
		}
	}
	return original
}
