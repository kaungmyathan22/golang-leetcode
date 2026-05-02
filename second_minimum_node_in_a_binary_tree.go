package main

func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	minVal := root.Val

	dfs(root, minVal, &ans)

	return ans
}

func dfs(node *TreeNode, minVal int, ans *int) {
	if node == nil {
		return
	}

	if node.Val > minVal {
		if *ans == -1 || node.Val < *ans {
			*ans = node.Val
		}
		return
	}

	dfs(node.Left, minVal, ans)
	dfs(node.Right, minVal, ans)
}
