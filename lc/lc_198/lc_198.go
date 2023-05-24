package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1}
	res := rob(arr)
	fmt.Println(res)
}

func rob(nums []int) int {
	//dp[i]表示偷到第i家的最大收获
	// dp[i] = max(dp[i-1], dp[i-2]+num[i-2])
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
