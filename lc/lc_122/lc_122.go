package main

func maxProfit(prices []int) int {
	count := 0
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i] > prices[i-1] {
			count += (prices[i] - prices[i-1])
		}
	}
	return count
}
