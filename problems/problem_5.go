package problems

import (
	"math"
)

func TotalMoney(n int) int {
	total := 0
	increment := 1
	totalWeeks := int(math.Ceil(float64(n) / float64(7)))
	for i := 0; i < totalWeeks; i++ {
		innerControl := func() int {
			if n > 7 {
				return n - i*7
			}
			return n
		}()
		if innerControl > 7 {
			innerControl = 7
		}
		for ii := 0; ii < innerControl; ii++ {
			total += ii + increment
		}
		increment++
	}
	return total
}
