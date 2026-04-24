package main

func isTrionic(nums []int) bool {
	n := len(nums)
	// Need at least 3 elements to have valid p and q indices.
	if n < 3 {
		return false
	}

	// p must be between 1 and n-3 (because q > p and q ≤ n-2)
	for p := 1; p <= n-3; p++ {
		// q must be between p+1 and n-2
		for q := p + 1; q <= n-2; q++ {
			// Check strictly increasing from 0 to p
			inc1 := true
			for i := 0; i < p; i++ {
				if nums[i] >= nums[i+1] {
					inc1 = false
					break
				}
			}
			if !inc1 {
				continue
			}

			// Check strictly decreasing from p to q
			dec := true
			for i := p; i < q; i++ {
				if nums[i] <= nums[i+1] {
					dec = false
					break
				}
			}
			if !dec {
				continue
			}

			// Check strictly increasing from q to n-1
			inc2 := true
			for i := q; i < n-1; i++ {
				if nums[i] >= nums[i+1] {
					inc2 = false
					break
				}
			}
			if inc2 {
				return true
			}
		}
	}
	return false
}
