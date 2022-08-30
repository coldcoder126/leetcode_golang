package main

func main() {

}

// 尽量让石头分成两份相同大小的堆
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true
	// i是石头，j是目标
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j-stones[i]] {
				dp[j] = true
			}
		}
	}
	if dp[target] && sum%2 == 0 {
		return 0
	}
	for i := target; i > 0; i-- {
		if dp[i] {
			return sum - i*2
		}
	}
	return stones[0]
}
