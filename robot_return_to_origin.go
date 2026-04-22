package main

func robotReturnToOrigin(moves string) bool {
	horizontal := 0
	vertical := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			vertical++
		}
		if moves[i] == 'D' {
			vertical--
		}
		if moves[i] == 'L' {
			horizontal--
		}
		if moves[i] == 'R' {
			horizontal++
		}
	}
	return horizontal == 0 && vertical == 0
}