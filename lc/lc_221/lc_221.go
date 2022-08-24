package main

import "fmt"

func main() {
	arr := make([][]byte, 2)
	//arr[0] = []byte{1, 0, 1, 0, 0}
	//arr[1] = []byte{1, 0, 1, 1, 1}
	//arr[2] = []byte{1, 1, 1, 1, 1}
	//arr[3] = []byte{1, 0, 0, 1, 0}
	arr[0] = []byte{0, 1}
	arr[1] = []byte{1, 0}

	res := maximalSquare(arr)
	fmt.Println(res)
}

func maximalSquare(matrix [][]byte) int {
	l := len(matrix)
	w := len(matrix[0])
	dp := make([][]int, l)
	maxEdge := 0
	for i := 0; i < l; i++ {
		dp[i] = make([]int, w)
		if matrix[i][0] == 0 {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
			maxEdge = 1
		}
	}
	for j := 0; j < w; j++ {
		if matrix[0][j] == 0 {
			dp[0][j] = 0
		} else {
			dp[0][j] = 1
			maxEdge = 1
		}
	}

	for i := 1; i < l; i++ {
		for j := 1; j < w; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				val := min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
				dp[i][j] = val + 1
				maxEdge = max(maxEdge, val+1)
			}
		}
	}
	return maxEdge * maxEdge
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
