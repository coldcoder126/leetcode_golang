package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, -3, -4, 5, -2, 5, 0, 6, 9}
	//arr := []int{2, 3, -2, 4}
	res := maxProduct(arr)

	fmt.Println(res)
}

func maxProduct(nums []int) int {

	l := len(nums)
	maxPos := make([]int, l)
	minNeg := make([]int, l)
	maxPos[0] = nums[0]
	minNeg[0] = nums[0]
	maxNum := nums[0]
	for i := 1; i < l; i++ {
		// 三选一：max(nums[i], nums[i]*maxPos[i-1], nums[i]*minNeg[i-1] )
		maxPos[i] = max(max(nums[i], nums[i]*maxPos[i-1]), minNeg[i-1]*nums[i])
		minNeg[i] = min(min(nums[i], nums[i]*minNeg[i-1]), maxPos[i-1]*nums[i])
		maxNum = max(max(maxPos[i], minNeg[i]), maxNum)
	}
	return maxNum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
