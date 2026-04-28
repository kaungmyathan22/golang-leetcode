package main

func missingNumber(nums []int) int {
	n := len(nums)
	return n*(n+1)/2 - s(nums)
}

func s(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
