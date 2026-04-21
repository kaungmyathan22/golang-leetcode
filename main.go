package main

import "fmt"

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
	// fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: nil, Right: nil}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
	// fmt.Println(InorderTraversal(&TreeNode{Val: 1, Left: nil, Right: nil}))
	// fmt.Println(isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: nil}, &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}))
	// fmt.Println(IsSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	// fmt.Println(IsSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}))
	// fmt.Println(ReverseInteger(123))
	// fmt.Println(ReverseInteger(-123))
	// fmt.Println(ReverseInteger(1534236469))
	// fmt.Println(ReverseInteger(-2147483648))
	// fmt.Println(ReverseInteger(120))
	// fmt.Println(StringToInteger("+1")) // 1
	// fmt.Println(StringToInteger("42"))
	// fmt.Println(StringToInteger("   -042"))
	// fmt.Println(StringToInteger("-42")) // -42
	// fmt.Println(StringToInteger("4193 with words"))
	// fmt.Println(StringToInteger("words and 987"))
	// fmt.Println(StringToInteger("-91283472332")) // -2147483648
	// fmt.Println(IntegerToRoman(3))    // III
	// fmt.Println(IntegerToRoman(58))   // LVIII
	// fmt.Println(IntegerToRoman(1994)) // MCMXCIV
	// fmt.Println(IntegerToRoman(3749)) // MMMDCCXLIX
	// fmt.Println(GenerateParentheses(3))
	// fmt.Println(GenerateParentheses(1))
	// fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))              // 5
	// fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))                 // 0
	// fmt.Println(MaxProfit([]int{1, 2}))                          // 1
	// fmt.Println(MaxProfit([]int{2, 1, 2, 1, 0, 1, 2}))           // 1
	// fmt.Println(MaxProfit([]int{2, 4, 1}))                       // 2
	// fmt.Println(MaxProfit([]int{3, 2, 6, 5, 0, 3}))              // 4
	// fmt.Println(MaxProfit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 9
	// fmt.Println(MaxProfit([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})) // 0
	// fmt.Println(MaxProfit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 9
	// fmt.Println(minMirrorPairDistance([]int{12, 21, 45, 33, 54})) // 9
	// fmt.Println(minMirrorPairDistance([]int{10, 11, 12, 13}))     // 1
	// fmt.Println(minMirrorPairDistance([]int{120, 210}))           // 90
	// fmt.Println(minMirrorPairDistance([]int{21, 120})) // 90
	// fmt.Println(solveQueries([]int{1, 3, 1, 4, 1, 3, 2}, []int{0, 3, 5})) // [2, -1, 3]
	// fmt.Println(solveQueries([]int{1, 2, 3, 4}, []int{0, 1, 2, 3}))       // [-1,-1,-1,-1]
	// fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))     // 24
	// fmt.Println(findFinalValue([]int{2, 7, 9}, 4))            // 4
	// fmt.Println(findFinalValue([]int{8, 19, 4, 2, 15, 3}, 2)) // 16
	// fmt.Println(mirrorDistance(123)) // 123
	// fmt.Println(mirrorDistance(25))  // 123
	// fmt.Println(mirrorDistance(10))  // 123
	// fmt.Println(mirrorDistance(7))   // 123
	// fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 3, 2)) // 1
	// fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3)) // 1
	// fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})) // 7 -> 0 -> 8
	// fmt.Println(addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0}))
	// fmt.Println(addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}})) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
	// fmt.Println(ClosestTarget([]string{"apple", "orange", "banana", "pear", "pineapple"}, "pear", 2)) // 2
	// fmt.Println(ClosestTarget([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1))         // 2
	// fmt.Println(ClosestTarget([]string{"a", "b", "leetcode"}, "leetcode", 0))                         // 2
	// fmt.Println(ClosestTarget([]string{"i", "eat", "leetcode"}, "ate", 0))                            // 2
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1, 2, 1, 1, 3}))
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1, 1, 2, 3, 2, 1, 2}))
	// fmt.Println(MinimumDistanceBetweenThreeEqualElements([]int{1}))
	fmt.Println(MinimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}))
}
