package main

func maxSubArray(nums []int) int {
	preMax := 0
	maxVal := nums[0]
	for i := 0; i < len(nums); i++ {
		preMax = max(nums[i], preMax+nums[i])
		maxVal = max(maxVal, preMax)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
