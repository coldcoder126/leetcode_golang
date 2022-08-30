package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 1, 1, 1}
	res := findTargetSumWays(arr, 3)
	fmt.Println(res)
}

// 所有带元素 和为sum，所有带-号的元素，和为neg，target = (sum - neg) - neg
// neg = (sum - target)/2，如果sum-target不是非负偶数 ，直接返回0
// 转化为在元素中找一些元素组成neg

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (sum-target) < 0 || (sum-target)%2 != 0 {
		return 0
	}
	neg := (sum - target) / 2

	dp := make([]int, neg+1)

	dp[0] = 1
	// i是item；j是neg
	for i := 0; i < len(nums); i++ {
		for j := neg; j >= nums[i]; j-- {
			if dp[j-nums[i]] > 0 {
				dp[j] += 1
			}
		}
	}
	return dp[neg]
}
