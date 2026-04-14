package main

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(MaxDepth(root.Left), MaxDepth(root.Right))
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
