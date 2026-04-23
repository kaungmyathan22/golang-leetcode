package main

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	tmp := make([][]int, n)
	// Make a deep copy of mat so we do not modify the original
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, n)
		copy(tmp[i], mat[i])
	}
	for r := 0; r < 4; r++ {
		if isEqual(tmp, target) {
			return true
		}
		tmp = rotate90(tmp)
	}
	return false
}

func isEqual(a, b [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate90(mat [][]int) [][]int {
	n := len(mat)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-1-i] = mat[i][j]
		}
	}
	return res
}
