package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Given the root of a binary tree, return all root-to-leaf paths in any order.
 * A leaf is a node with no children.
 */

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	paths := make([]string, 0)
	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)
	for _, path := range leftPaths {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+path)
	}
	for _, path := range rightPaths {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+path)
	}
	return paths
}
