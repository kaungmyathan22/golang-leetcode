package main

func MinOperations(s string) int {
	n := len(s)
	// Two patterns: start with '0' or start with '1'
	countStartWith0 := 0
	countStartWith1 := 0

	for i := 0; i < n; i++ {
		expected0 := byte('0')
		expected1 := byte('1')
		if i%2 == 1 {
			expected0 = '1'
			expected1 = '0'
		}
		if s[i] != expected0 {
			countStartWith0++
		}
		if s[i] != expected1 {
			countStartWith1++
		}
	}
	if countStartWith0 < countStartWith1 {
		return countStartWith0
	}
	return countStartWith1
}
