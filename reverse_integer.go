package main

import "math"

func ReverseInteger(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	if x == math.MinInt32 {
		return 0
	}
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if sign*reversed > math.MaxInt32 {
		return 0
	}
	if sign*reversed < math.MinInt32 {
		return 0
	}
	return sign * reversed
}
