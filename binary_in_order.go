package main

func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)
	return result
}
