package main

func reverseSquareSubmatrix(matrix [][]int, x, y, k int) [][]int {
	for rowTop, rowBottom := x, x+k-1; rowTop < rowBottom; rowTop, rowBottom = rowTop+1, rowBottom-1 {
		for col := y; col < y+k; col++ {
			// Swap the elements in the current column between the top and bottom rows of the submatrix
			matrix[rowTop][col], matrix[rowBottom][col] = matrix[rowBottom][col], matrix[rowTop][col]
		}
	}
	return matrix
}
