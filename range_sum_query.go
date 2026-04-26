package main

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	NumArray := NumArray{
		prefixSum: []int{0},
	}
	for _, num := range nums {
		NumArray.prefixSum = append(NumArray.prefixSum, NumArray.prefixSum[len(NumArray.prefixSum)-1]+num)
	}
	return NumArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}
