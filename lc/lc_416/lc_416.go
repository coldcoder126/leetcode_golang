package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 2, 1, 1}
	res := canPartition2(arr)
	//res := withOneDim(arr)
	fmt.Println(res)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canPartition2(nums []int) bool {
	sum := 0
	//如果和为一个奇数，则不可能
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum/2 + 1
	//数组中是否有子集可以组成和为sum
	//这个问题转换为0-1背包问题，dp[j]表示容量为j的背包装的物品的重量即物品i的重量是num[i],值也是num[i],

	//递推公式： dp[j] = max(dp[j-i]) dp[i][j] = max(dp[i-1][j], dp[i-1][j-num[i]]+num[i])
	dp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		t := make([]int, sum)
		dp = append(dp, t)
	}

	//初始化
	for i := 0; i < sum; i++ {
		if i >= nums[0] {
			dp[0][i] = nums[0]
		}
	}
	//递推
	for i := 1; i < len(nums); i++ {
		for j := 0; j < sum; j++ {
			if j >= nums[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
		if dp[i][sum-1] == sum-1 {
			return true
		}
	}
	return false
}

// 用一维数组做dp
func withOneDim(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 || len(nums) == 1 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[j-nums[i]] {
				dp[j] = true
			}
		}
	}
	fmt.Println(dp)
	return dp[target]
}

/*
*
执行用时：20 ms, 在所有 Go 提交中击败了71.63%的用户
内存消耗：3.5 MB, 在所有 Go 提交中击败了51.66%的用户
*/
func withOneDim2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 || len(nums) == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
			if dp[j] == target {
				return true
			}
		}
	}
	return false
}
