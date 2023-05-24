package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	//1.将书按高度排序

	//dp[i] 表示第i本书可以的最低高度，当前这本书可以选择加入当前层书架或下一层书架
	//dp[i] =
	dp := make([]int, len(books)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < len(books); i++ {
		curWid := shelfWidth
		if curWid > books[i][0] {
			curWid = curWid - books[i][0]
			dp[i] = 0
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
