package main

func Generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		var arr = make([]int, i+1)
		arr[0], arr[i] = 1, 1
		for j := 1; j < i; j++ {
			arr[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = arr
	}
	return res
}
