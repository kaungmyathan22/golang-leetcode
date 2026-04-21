package main

func SpecialPositionsInABinaryMatrix(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])

	rowSum := make([]int, m)
	colSum := make([]int, n)

	// Calculate number of 1's in each row and column
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSum[i] += mat[i][j]
			colSum[j] += mat[i][j]
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rowSum[i] == 1 && colSum[j] == 1 {
				count++
			}
		}
	}
	return count
}
