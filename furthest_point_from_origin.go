package main

func furthestDistanceFromOrigin(moves string) int {
	left := 0
	right := 0
	for _, move := range moves {
		if move == 'L' {
			left++
		} else if move == 'R' {
			right++
		} else if move == '_' {
			continue
		}
	}
	underscoreCount := 0
	for _, move := range moves {
		if move == '_' {
			underscoreCount++
		}
	}
	dis := left - right
	if dis < 0 {
		dis = -dis
	}
	return dis + underscoreCount
}
