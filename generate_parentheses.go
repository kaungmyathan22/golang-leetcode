package main

import "fmt"

func GenerateParentheses(n int) []string {
	result := []string{}
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, n int, result *[]string) {
	fmt.Println("Current", current, "Open", open, "Close", close, "N", n)
	if open == close && open == n {
		fmt.Println("Adding to result", current)
		*result = append(*result, current)
		fmt.Println("Result", *result)
		return
	}
	if open < n {
		backtrack(current+"(", open+1, close, n, result)
	}
	if close < open {
		backtrack(current+")", open, close+1, n, result)
	}
}
