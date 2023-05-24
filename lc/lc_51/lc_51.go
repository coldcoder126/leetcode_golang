package main

import (
	"fmt"
	"strings"
)

func main() {
	res := solveNQueens(4)
	fmt.Println(res)
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	//ans := make([]string,0)
	m := make([][]int, 0)
	for i := 0; i < n; i++ {
		t := make([]int, n)
		m = append(m, t)
	}
	rowFlag := make([]int, n)

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			t := itos(m)
			res = append(res, t)
			return
		}

		for col := 0; col < n; col++ {
			if check(m, row, col) {
				m[row][col] = 1
				rowFlag[row] = 1
				dfs(row + 1)
				rowFlag[row] = 0
				m[row][col] = 0
			}
		}
	}
	dfs(0)
	return res
}

// 检查当前matrix中，在[row,col]位置的会不会和其他有冲突，只需要检查正上、左上和右上
func check(m [][]int, row, col int) bool {
	n := len(m)
	//检查同一行和同一列
	for i := row; i >= 0; i-- {
		if m[i][col] == 1 {
			return false
		}
		if row-i >= 0 && col-i >= 0 && m[row-i][col-i] == 1 {
			return false
		}
		if row-i >= 0 && col+i < n && m[row-i][col+i] == 1 {
			return false
		}
	}
	return true
}

func itos(m [][]int) (s []string) {
	var builder strings.Builder
	for i := 0; i < len(m); i++ {
		for _, j := range m[i] {
			if j == 0 {
				builder.WriteString(".")
			} else {
				builder.WriteString("Q")
			}
		}
		s = append(s, builder.String())
		builder.Reset()
	}
	return s
}
