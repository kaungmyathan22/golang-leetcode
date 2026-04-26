package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n < 0 {
		return false
	}
	for i := 1; i < n; i++ {
		result := i * i * i
		fmt.Println(result, "result", i, "i", n, "n")
		if result == n {
			return true
		}
		if result > n {
			return false
		}
	}
	return false
}
