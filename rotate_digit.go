package main

func rotateDigit(n int) int {
	rot := map[int]int{
		0: 0,
		1: 1,
		2: 5,
		5: 2,
		6: 9,
		8: 8,
		9: 6,
	}

	count := 0

	for i := 1; i <= n; i++ {
		if isGood(i, rot) {
			count++
		}
	}

	return count
}

func isGood(num int, rot map[int]int) bool {
	changed := false
	original := num

	for num > 0 {
		digit := num % 10

		val, ok := rot[digit]
		if !ok {
			return false
		}

		if val != digit {
			changed = true
		}

		num /= 10
	}

	return changed && original > 0
}
