package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	total := 0

	for i := 0; i < len(timeSeries)-1; i++ {
		gap := timeSeries[i+1] - timeSeries[i]

		if gap < duration {
			total += gap
		} else {
			total += duration
		}
	}

	return total + duration
}
