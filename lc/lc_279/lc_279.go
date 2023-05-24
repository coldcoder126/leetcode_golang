package main

import (
	"fmt"
	"math"
)

func main() {
	x := 18
	res := numSquares(x)
	fmt.Println(res)
}

func numSquares(n int) int {
	// dp[j]表示值为i的完全平方数的最少数量
	// dp[j] = min(dp[j-i*i]+1,dp[j])
	dp := make([]int, n+1)
	maxNum := int(math.Sqrt(float64(n)))

	//初始化
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= maxNum; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
