package main

import "fmt"

var res [][]int
var temp []int

func main() {
	combine(4, 2)
}
func combine(n int, k int) [][]int {
	dfs(1, n, k)
	fmt.Println(res)
	return res
}

// start 开始的位置， deep 深度 ， k 给的深度
func dfs(start, n, k int) {
	if len(temp) == k {
		t := make([]int, k)
		copy(t, temp)
		//copy(t, temp)
		res = append(res, t)
		return
	}
	//temp = append(temp, start)
	for i := start; i <= n; i++ {
		temp = append(temp, i)
		dfs(i+1, n, k)
		temp = temp[:len(temp)-1]
	}
}
