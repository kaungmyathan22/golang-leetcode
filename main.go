package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{1, 1, 2}
	// (removeDuplicates(nums1))
	// nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// (removeDuplicates(nums2))
	// fmt.Println(nums2)
	// fmt.Println(problems.TotalMoney(4))
	// fmt.Println(problems.TotalMoney(10))
	// fmt.Println(problems.TotalMoney(20))
	// fmt.Println(problems.LargestOddNumber("52"))
	// fmt.Println(problems.LargestOddNumber("4206"))
	// fmt.Println(problems.LargestOddNumber("35427"))
	// fmt.Println(problems.LargestOddNumber("10133890"))
	// fmt.Println(problems.LargestOddNumber("239537672423884969653287101"))
	// fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	// fmt.Println(findWords([]string{"omk"}))
	// fmt.Println(findWords([]string{"adsdf", "sfd"}))
	fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: nil, Right: nil}))
}
