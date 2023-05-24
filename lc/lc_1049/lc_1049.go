package main

import "fmt"

func main() {
	arr := []int{31, 26, 33, 21, 40}
	res := lastStoneWeightII(arr)
	fmt.Println(res)

}

// 尽量让石头分成两份相同大小的堆
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	target := sum / 2
	//dp[i]的含义是容量为i的背包可以装的最大值
	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	//sum可能为奇数也可能为偶数，但是不管哪个都是 (sum-dp[target])-dp[target]
	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
