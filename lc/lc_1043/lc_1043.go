package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	res := maxSumAfterPartitioning(arr, 4)
	fmt.Println(res)

}

func maxSumAfterPartitioning(arr []int, k int) int {
	//dp[j]表示 截止到第j个位置的最大值 dp[1]
	//dp[j] = max((dp[j-1]+arr[j]), ... , dp[j-k]+max(arr[j-k:j+1])*k )
	dp := make([]int, len(arr)+1)
	dp[1] = arr[0]

	for i := 1; i < len(arr); i++ {
		for j := 1; j <= k; j++ { //j是窗口大小
			if i-j+1 >= 0 {
				dp[i+1] = max(dp[i+1], (maxInArr(arr[i-j+1:i+1])*j + dp[i-j+1]))
			}
		}
	}
	return dp[len(arr)]

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxInArr(arr []int) int {
	res := math.MinInt
	for i := 0; i < len(arr); i++ {
		res = max(arr[i], res)
	}
	return res
}
