package main

func findTheWinner(n int, k int) int {
	winner := 0
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1
}

// func main() {
// 	fmt.Println(findTheWinner(5, 2))
// 	fmt.Println(findTheWinner(6, 5))
// }
