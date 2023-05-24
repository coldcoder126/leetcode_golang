package main

import "fmt"

func main() {
	arr := []int{1, 2, 5}
	res := change(5, arr)
	fmt.Println(res)
}

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	//dp[i]表示可以凑成总金额i的有几种方式
	//dp[j] = dp[j-coins[i]] + dp[j]
	dp := make([]int, amount+1)
	//初始化

	for i := 0; i <= amount; i++ {
		if i >= coins[0] && i%coins[0] == 0 {
			dp[i] = 1
			dp[0] = 1
		}
	}

	for i := 1; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] = dp[j-coins[i]] + dp[j]
			}
		}
	}
	return dp[amount]
}
