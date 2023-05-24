package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{83, 20, 17, 43, 52, 78, 68, 45}
	res := longestArithSeqLength(arr)
	fmt.Println(res)
}

func longestArithSeqLength(nums []int) int {
	minNum, maxNum := math.MaxInt32, math.MinInt
	for i := 0; i < len(nums); i++ {
		minNum = min(minNum, nums[i])
		maxNum = max(maxNum, nums[i])
	}
	ans := 0
	item := maxNum - minNum

	//dp[i][j]表示在num[i]以j为公差的最大数组的长度
	//dp[i][j] = max(dp[i][j],dp[k][j]+1)；其中num[k]=num[i]-j
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2*item+1)
	}

	//遍历可能的公差
	for i := 1; i < len(nums); i++ {
		for k := 0; k < i; k++ {
			j := nums[i] - nums[k] + item
			dp[i][j] = max(dp[i][j], dp[k][j]+1)
			ans = max(ans, dp[i][j])
		}
	}
	return ans + 1
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
