package main

import "fmt"

func main() {
	arr := make([][]byte, 4)
	arr[0] = []byte{1, 0, 1, 1, 1}
	arr[1] = []byte{1, 0, 1, 0, 1}
	arr[2] = []byte{1, 1, 1, 0, 1}
	arr[3] = []byte{0, 0, 0, 0, 0}
	x := numIslands(arr)
	fmt.Println(x)
}

func numIslands(grid [][]byte) int {
	l := len(grid)
	w := len(grid[0])
	mark := make([][]int, l)
	for i := 0; i < l; i++ {
		mark[i] = make([]int, w)
	}

	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == 1 && mark[i][j] == 0 {
			mark[i][j] = 1
			if i+1 < l {
				dfs(i+1, j) //向下
			}
			if j+1 < w {
				dfs(i, j+1) //向右
			}
			if j > 0 {
				dfs(i, j-1) //向左
			}
			if i > 0 {
				dfs(i-1, j) //向上
			}
		}
		return
	}

	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if mark[i][j] == 0 { // 未访问过
				if grid[i][j] == 0 {
					mark[i][j] = 1
					continue
				} else {
					dfs(i, j)
					count++
				}
			}
		}
	}
	return count
}
