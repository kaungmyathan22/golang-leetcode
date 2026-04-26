package main

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i/2]
		if i%2 == 1 {
			ans[i]++
		}
	}

	return ans
}
