package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 0, 3, 2, 3}
	res := lengthOfLIS(arr)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
