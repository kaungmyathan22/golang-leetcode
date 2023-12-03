package main

func removeElement(nums []int, val int) int {
	j := 0
	for index, value := range nums {
		if value != val {
			nums[j], nums[index] = nums[index], nums[j]
			j++
		}
	}
	return j
}

// func removeElement(nums []int, val int) int {
// 	j := 0
// 	for index:=0 ; index < len(nums); index++ {
// 		if nums[index] == val {
// 			nums[j], nums[index] = nums[index], nums[j]
// 			j++
// 		}
// 	}
// 	return j
// }

// func removeElement(nums []int, val int) int {
// 	j := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != val {
// 			nums[j], nums[i] = nums[i], nums[j]
// 			j++
// 		}
// 	}
// 	return j
// }
