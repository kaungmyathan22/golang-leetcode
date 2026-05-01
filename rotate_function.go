package main

func maxRotateFunction(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	total := 0
	f := 0

	for i, v := range nums {
		total += v
		f += i * v
	}

	ans := f

	for k := 1; k < n; k++ {
		f = f + total - n*nums[n-k]

		if f > ans {
			ans = f
		}
	}

	return ans
}
