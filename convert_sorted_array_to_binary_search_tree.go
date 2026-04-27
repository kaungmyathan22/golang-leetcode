package main

import "fmt"

func debug(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	debug(node.Left)
	debug(node.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
