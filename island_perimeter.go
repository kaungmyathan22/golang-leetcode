package main

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	shared := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count += 1
				if i > 0 && grid[i-1][j] == 1 {
					shared += 1
				}
				if j > 0 && grid[i][j-1] == 1 {
					shared += 1
				}
			}
		}
	}
	return count*4 - shared*2
}
