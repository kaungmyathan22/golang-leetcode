package main

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	NumArray := NumArray{
		// префиксный массив, тоесть если например начальный массив nums = [1,2,3] то он станет -> [0,1,3,6]
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
