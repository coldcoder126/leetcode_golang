package main

import "fmt"

func main() {
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r := len(obstacleGrid)
	c := len(obstacleGrid[0])
	fmt.Println(r, c)
	dp := make([][]int, 0)
	for i := 0; i < r; i++ {
		t := make([]int, c)
		dp = append(dp, t)
	}
	//初始化一些状态
	for i := 0; i < r; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < c; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	//dp[0][j]=1 ; dp[i][0]=1 ; dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[r-1][c-1]
}
