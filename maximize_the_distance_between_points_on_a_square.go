package main

import (
	"sort"
)

func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	perim := 4 * side

	// Step 1: map to 1D perimeter
	pos := make([]int, n)
	for i, p := range points {
		x, y := p[0], p[1]

		if x == 0 {
			pos[i] = y
		} else if y == side {
			pos[i] = side + x
		} else if x == side {
			pos[i] = 3*side - y
		} else { // y == 0
			pos[i] = 4*side - x
		}
	}

	sort.Ints(pos)

	// duplicate for circular handling
	ext := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ext[i] = pos[i]
		ext[i+n] = pos[i] + perim
	}

	// check feasibility
	can := func(d int) bool {
		for start := 0; start < n; start++ {
			count := 1
			last := ext[start]
			for count < k {
				// binary search next valid point
				target := last + d
				j := sort.Search(len(ext), func(i int) bool {
					return ext[i] >= target
				})

				if j >= start+n {
					break
				}

				last = ext[j]
				count++
			}

			if count == k {
				// also check circular wrap constraint
				if last-ext[start] <= perim-d {
					return true
				}
			}
		}
		return false
	}

	// binary search answer
	low, high := 0, perim
	ans := 0

	for low <= high {
		mid := (low + high) / 2
		if can(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
