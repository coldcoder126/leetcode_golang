package main

import "math"

func main() {
}

func maxProfit(prices []int) int {
	preMin := math.MaxInt
	maxVal := 0
	for i := 0; i < len(prices); i++ {
		preMin = min(preMin, prices[i])
		maxVal = max(prices[i]-preMin, maxVal)
	}
	return maxVal
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
