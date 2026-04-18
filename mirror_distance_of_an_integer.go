package main

import "math"

func mirrorDistance(n int) int {
	return int(math.Abs(float64(reverse_(n) - n)))
}

func reverse_(x int) int {
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return reversed
}
