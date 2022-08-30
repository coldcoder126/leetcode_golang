package main

func integerBreak(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	for i := 4; i <= n; i++ {
		dp[i] = max(2*max(dp[i-2], i-2), 3*max(dp[i-3], i-3))
	}
	return dp[n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
