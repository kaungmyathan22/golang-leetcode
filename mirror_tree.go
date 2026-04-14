package main

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Compare(root.Left, root.Right)
}

func Compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && Compare(left.Left, right.Right) && Compare(left.Right, right.Left)
}
