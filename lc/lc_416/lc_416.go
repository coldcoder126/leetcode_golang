package main

import "fmt"

func main() {
	arr := []int{2, 2, 1, 1}
	//res := canPartition(arr)
	res := withOneDim(arr)
	fmt.Println(res)

}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	//sort.Ints(nums)

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}
	//第一行先特殊处理一下
	if nums[0] <= target {
		dp[0][nums[0]] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			// 将上一行的状态取下来
			if dp[i-1][j] == 1 {
				dp[i][j] = dp[i-1][j]
			}
		}
		// 放入，要求放入之后，不能超过背包的总重量（target）
		// 上一步的所有状态+当前物品的
		for j := 0; j <= target-nums[i]; j++ {
			if dp[i-1][j] == 1 {
				dp[i][j+nums[i]] = 1
			}
		}
	}
	fmt.Println(dp)
	for i := 0; i < len(nums); i++ {
		if dp[i][target] == 1 {
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
