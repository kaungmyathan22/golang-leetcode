package main

func majorityElement(nums []int) int {
	dir := map[int]int{}
	threadshold := len(nums) / 3
	for _, v := range nums {
		dir[v]++
	}

	most_occurance_number := 1
	most_occurance_count := 1
	for key, value := range dir {
		if value > most_occurance_count && value > threadshold {
			most_occurance_number = key
		}
	}
	return most_occurance_number
}

/**
func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
*/
