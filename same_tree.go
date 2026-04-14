package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil || p == nil {
		return q == p
	}
	if q.Val != p.Val {
		return false
	}
	isLeftSame := isSameTree(p.Left, q.Left)
	isRightSame := isSameTree(p.Right, q.Right)
	return isLeftSame && isRightSame
}
