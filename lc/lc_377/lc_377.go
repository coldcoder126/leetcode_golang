package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	res := combinationSum4(arr, 4)
	fmt.Println(res)

}

func combinationSum4(nums []int, target int) int {

	//dp[j]表示和为j的组合数
	//递推：dp[j] = dp[j]+dp[j-nums[i]]
	dp := make([]int, target+1)

	//初始化
	for i := 0; i <= target; i++ {
		if i%nums[0] == 0 {
			dp[i] = 1
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}