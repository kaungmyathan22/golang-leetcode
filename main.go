package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	MergeSortedArray(nums1, 3, nums2, 3)
	fmt.Println(nums1)
	nums1 = []int{0}
	nums2 = []int{1}
	MergeSortedArray(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}
